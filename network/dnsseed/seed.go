package dnsseed

import (
	"context"
	"fmt"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/mconfig/dagconfig"
	"github.com/Metchain/Metblock/mconfig/externalapi"
	internalpeer "github.com/Metchain/Metblock/protoserver/peer"
	"github.com/Metchain/Metblock/utils/mstime"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type OnSeed func(addrs []*appmessage.NetAddress)

type LookupFunc func(string) ([]net.IP, error)

func SeedFromDNS(dagParams *dagconfig.Params, customSeed string, includeAllSubnetworks bool,
	subnetworkID *externalapi.DomainSubnetworkID, lookupFn LookupFunc, seedFn OnSeed) {

	var dnsSeeds []string
	if customSeed != "" {
		dnsSeeds = []string{customSeed}
	} else {
		dnsSeeds = dagParams.DNSSeeds
	}

	for _, dnsseed := range dnsSeeds {
		host := dnsseed

		if !includeAllSubnetworks {
			if subnetworkID != nil {
				host = fmt.Sprintf("%c%s.%s", SubnetworkIDPrefixChar, subnetworkID, host)
			} else {
				host = fmt.Sprintf("%c.%s", SubnetworkIDPrefixChar, host)
			}
		}

		spawn("SeedFromDNS", func() {
			randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

			seedPeers, err := lookupFn(host)
			if err != nil {
				log.Infof("DNS discovery failed on seed %s: %s", host, err)
				return
			}
			numPeers := len(seedPeers)

			log.Infof("%d addresses found from DNS seed %s", numPeers, host)

			if numPeers == 0 {
				return
			}
			addresses := make([]*appmessage.NetAddress, len(seedPeers))
			// if this errors then we have *real* problems
			intPort, _ := strconv.Atoi(dagParams.DefaultPort)
			for i, peer := range seedPeers {
				addresses[i] = appmessage.NewNetAddressTimestamp(
					// seed with addresses from a time randomly selected
					// between 3 and 7 days ago.
					mstime.Now().Add(-1*time.Second*time.Duration(secondsIn3Days+
						randSource.Int31n(secondsIn4Days))),
					peer, uint16(intPort))
			}

			seedFn(addresses)
		})
	}
}

// SeedFromGRPC send gRPC request to get list of peers for a given host
func SeedFromGRPC(dagParams *dagconfig.Params, customSeed string, includeAllSubnetworks bool,
	subnetworkID *externalapi.DomainSubnetworkID, seedFn OnSeed) {

	var grpcSeeds []string
	if customSeed != "" {
		grpcSeeds = []string{customSeed}
	} else {
		grpcSeeds = dagParams.GRPCSeeds
	}

	for _, host := range grpcSeeds {
		spawn("SeedFromGRPC", func() {
			randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

			conn, err := grpc.Dial(host, grpc.WithInsecure())
			client := internalpeer.NewPeerServiceClient(conn)
			if err != nil {
				log.Warnf("Failed to connect to gRPC server: %s", host)
			}

			var subnetID []byte
			if subnetworkID != nil {
				subnetID = subnetworkID[:]
			} else {
				subnetID = nil
			}

			req := &internalpeer.GetPeersListRequest{
				SubnetworkID:          subnetID,
				IncludeAllSubnetworks: includeAllSubnetworks,
			}
			res, err := client.GetPeersList(context.Background(), req)

			if err != nil {
				log.Infof("gRPC request to get peers failed (host=%s): %s", host, err)
				return
			}

			seedPeers := fromProtobufAddresses(res.Addresses)

			numPeers := len(seedPeers)

			log.Infof("%d addresses found from DNS seed %s", numPeers, host)

			if numPeers == 0 {
				return
			}
			addresses := make([]*appmessage.NetAddress, len(seedPeers))
			// if this errors then we have *real* problems
			intPort, _ := strconv.Atoi(dagParams.DefaultPort)
			for i, peer := range seedPeers {
				addresses[i] = appmessage.NewNetAddressTimestamp(
					// seed with addresses from a time randomly selected
					// between 3 and 7 days ago.
					mstime.Now().Add(-1*time.Second*time.Duration(secondsIn3Days+
						randSource.Int31n(secondsIn4Days))),
					peer, uint16(intPort))
			}

			seedFn(addresses)
		})
	}
}

func fromProtobufAddresses(proto []*internalpeer.NetAddress) []net.IP {
	var addresses []net.IP

	for _, pbAddr := range proto {
		addresses = append(addresses, net.IP(pbAddr.IP))
	}

	return addresses
}
