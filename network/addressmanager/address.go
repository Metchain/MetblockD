package addressmanager

import (
	"encoding/binary"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/utils/mstime"
	"github.com/pkg/errors"
	"net"
	"strconv"
	"sync"
)

type AddressManager struct {
	store          *addressStore
	localAddresses *localAddressManager
	mutex          sync.Mutex
	cfg            *Config
}

// addressKey represents a pair of IP and port, the IP is always in V6 representation
type addressKey struct {
	port    uint16
	address ipv6
}

type ipv6 [net.IPv6len]byte

type address struct {
	netAddress            *appmessage.NetAddress
	connectionFailedCount uint64
}

func (am *AddressManager) AddAddresses(addresses ...*appmessage.NetAddress) error {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	for _, address := range addresses {
		err := am.addAddressNoLock(address)
		if err != nil {
			return err
		}
	}
	return nil
}

func (am *AddressManager) addAddressNoLock(netAddress *appmessage.NetAddress) error {
	if !IsRoutable(netAddress, am.cfg.AcceptUnroutable) {
		return nil
	}

	key := netAddressKey(netAddress)
	// We mark `connectionFailedCount` as 0 only after first success
	address := &address{netAddress: netAddress, connectionFailedCount: 1}
	err := am.store.add(key, address)
	if err != nil {
		return err
	}

	if am.store.notBannedCount() > maxAddresses {
		allAddresses := am.store.getAllNotBanned()

		maxConnectionFailedCount := uint64(0)
		toRemove := allAddresses[0]
		for _, address := range allAddresses[1:] {
			if address.connectionFailedCount > maxConnectionFailedCount {
				maxConnectionFailedCount = address.connectionFailedCount
				toRemove = address
			}
		}

		err := am.removeAddressNoLock(toRemove.netAddress)
		if err != nil {
			return err
		}
	}
	return nil
}

func New(cfg *Config, database database.Database) (*AddressManager, error) {
	addressStore, err := newAddressStore(database)
	if err != nil {
		return nil, err
	}
	localAddresses, err := newLocalAddressManager(cfg)
	if err != nil {
		return nil, err
	}

	return &AddressManager{
		store:          addressStore,
		localAddresses: localAddresses,
		cfg:            cfg,
	}, nil
}

func newAddressStore(database database.Database) (*addressStore, error) {
	addressStore := &addressStore{
		database:           database,
		notBannedAddresses: map[addressKey]*address{},
		bannedAddresses:    map[ipv6]*address{},
	}
	err := addressStore.restoreNotBannedAddresses()
	if err != nil {
		return nil, err
	}
	err = addressStore.restoreBannedAddresses()
	if err != nil {
		return nil, err
	}

	log.Infof("Loaded %d addresses and %d banned addresses",
		len(addressStore.notBannedAddresses), len(addressStore.bannedAddresses))

	return addressStore, nil
}

func newLocalAddressManager(cfg *Config) (*localAddressManager, error) {
	localAddressManager := localAddressManager{
		localAddresses: map[addressKey]*localAddress{},
		cfg:            cfg,
		lookupFunc:     cfg.Lookup,
	}

	err := localAddressManager.initListeners()
	if err != nil {
		return nil, err
	}

	return &localAddressManager, nil
}

// initListeners initializes the configured net listeners and adds any bound
// addresses to the address manager
func (lam *localAddressManager) initListeners() error {
	if len(lam.cfg.ExternalIPs) != 0 {
		defaultPort, err := strconv.ParseUint(lam.cfg.DefaultPort, 10, 16)
		if err != nil {
			log.Errorf("Can not parse default port %s for active DAG: %s",
				lam.cfg.DefaultPort, err)
			return err
		}

		for _, sip := range lam.cfg.ExternalIPs {
			eport := uint16(defaultPort)
			host, portstr, err := net.SplitHostPort(sip)
			if err != nil {
				// no port, use default.
				host = sip
			} else {
				port, err := strconv.ParseUint(portstr, 10, 16)
				if err != nil {
					log.Warnf("Can not parse port from %s for "+
						"externalip: %s", sip, err)
					continue
				}
				eport = uint16(port)
			}
			na, err := lam.hostToNetAddress(host, eport)
			if err != nil {
				log.Warnf("Not adding %s as externalip: %s", sip, err)
				continue
			}

			err = lam.addLocalNetAddress(na, ManualPrio)
			if err != nil {
				log.Warnf("Skipping specified external IP: %s", err)
			}
		}
	} else {
		// Listen for TCP connections at the configured addresses
		netAddrs, err := parseListeners(lam.cfg.Listeners)
		if err != nil {
			return err
		}

		// Add bound addresses to address manager to be advertised to peers.
		for _, addr := range netAddrs {
			listener, err := net.Listen(addr.Network(), addr.String())
			if err != nil {
				log.Warnf("Can't listen on %s: %s", addr, err)
				continue
			}
			addr := listener.Addr().String()
			err = listener.Close()
			if err != nil {
				return err
			}
			err = lam.addLocalAddress(addr)
			if err != nil {
				log.Warnf("Skipping bound address %s: %s", addr, err)
			}
		}
	}

	return nil
}

// hostToNetAddress returns a netaddress given a host address. If
// the host is not an IP address it will be resolved.
func (lam *localAddressManager) hostToNetAddress(host string, port uint16) (*appmessage.NetAddress, error) {
	ip := net.ParseIP(host)
	if ip == nil {
		ips, err := lam.lookupFunc(host)
		if err != nil {
			return nil, err
		}
		if len(ips) == 0 {
			return nil, errors.Errorf("no addresses found for %s", host)
		}
		ip = ips[0]
	}

	return appmessage.NewNetAddressIPPort(ip, port), nil
}

// addLocalNetAddress adds netAddress to the list of known local addresses to advertise
// with the given priority.
func (lam *localAddressManager) addLocalNetAddress(netAddress *appmessage.NetAddress, priority AddressPriority) error {
	if !IsRoutable(netAddress, lam.cfg.AcceptUnroutable) {
		return errors.Errorf("address %s is not routable", netAddress.IP)
	}

	lam.mutex.Lock()
	defer lam.mutex.Unlock()

	addressKey := netAddressKey(netAddress)
	address, ok := lam.localAddresses[addressKey]
	if !ok || address.score < priority {
		if ok {
			address.score = priority + 1
		} else {
			lam.localAddresses[addressKey] = &localAddress{
				netAddress: netAddress,
				score:      priority,
			}
		}
	}
	return nil
}

// addLocalAddress adds an address that this node is listening on to the
// address manager so that it may be relayed to peers.
func (lam *localAddressManager) addLocalAddress(addr string) error {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return err
	}

	if ip := net.ParseIP(host); ip != nil && ip.IsUnspecified() {
		// If bound to unspecified address, advertise all local interfaces
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return err
		}

		for _, addr := range addrs {
			ifaceIP, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				continue
			}

			// If bound to 0.0.0.0, do not add IPv6 interfaces and if bound to
			// ::, do not add IPv4 interfaces.
			if (ip.To4() == nil) != (ifaceIP.To4() == nil) {
				continue
			}

			netAddr := appmessage.NewNetAddressIPPort(ifaceIP, uint16(port))
			lam.addLocalNetAddress(netAddr, BoundPrio)
		}
	} else {
		netAddr, err := lam.hostToNetAddress(host, uint16(port))
		if err != nil {
			return err
		}

		lam.addLocalNetAddress(netAddr, BoundPrio)
	}

	return nil
}

func (as *addressStore) restoreNotBannedAddresses() error {
	cursor, err := as.database.Cursor(notBannedAddressBucket)
	if err != nil {
		return err
	}
	defer cursor.Close()
	for ok := cursor.First(); ok; ok = cursor.Next() {
		databaseKey, err := cursor.Key()
		if err != nil {
			return err
		}
		serializedKey := databaseKey.Suffix()
		key := as.deserializeAddressKey(serializedKey)

		serializedNetAddress, err := cursor.Value()
		if err != nil {
			return err
		}
		netAddress := as.deserializeAddress(serializedNetAddress)
		as.notBannedAddresses[key] = netAddress
	}
	return nil
}

func (as *addressStore) deserializeAddressKey(serializedKey []byte) addressKey {
	var ip ipv6
	copy(ip[:], serializedKey[:])

	port := binary.LittleEndian.Uint16(serializedKey[16:])

	return addressKey{
		port:    port,
		address: ip,
	}
}

func (as *addressStore) deserializeAddress(serializedAddress []byte) *address {
	ip := make(net.IP, 16)
	copy(ip[:], serializedAddress[:])

	port := binary.LittleEndian.Uint16(serializedAddress[16:])
	timestamp := mstime.UnixMilliseconds(int64(binary.LittleEndian.Uint64(serializedAddress[18:])))
	connectionFailedCount := binary.LittleEndian.Uint64(serializedAddress[26:])

	return &address{
		netAddress: &appmessage.NetAddress{
			IP:        ip,
			Port:      port,
			Timestamp: timestamp,
		},
		connectionFailedCount: connectionFailedCount,
	}
}

func (as *addressStore) restoreBannedAddresses() error {
	cursor, err := as.database.Cursor(bannedAddressBucket)
	if err != nil {
		return err
	}
	defer cursor.Close()
	for ok := cursor.First(); ok; ok = cursor.Next() {
		databaseKey, err := cursor.Key()
		if err != nil {
			return err
		}
		var ipv6 ipv6
		copy(ipv6[:], databaseKey.Suffix())

		serializedNetAddress, err := cursor.Value()
		if err != nil {
			return err
		}
		netAddress := as.deserializeAddress(serializedNetAddress)
		as.bannedAddresses[ipv6] = netAddress
	}
	return nil
}
