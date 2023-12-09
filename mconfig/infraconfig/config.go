package infraconfig

import (
	"fmt"
	"github.com/Metchain/Metblock/mconfig/externalapi"
	"github.com/Metchain/Metblock/mconfig/network"
	"github.com/Metchain/Metblock/mconfig/util"
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/Metchain/Metblock/version"
	"github.com/btcsuite/go-socks/socks"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func LoadConfig() (*Config, error) {
	cfgFlags := defaultFlags()

	// Pre-parse the command line options to see if an alternative config
	// file or the version flag was specified. Any errors aside from the
	// help message error can be ignored here since they will be caught by
	// the final parse below.
	preCfg := cfgFlags
	preParser := newConfigParser(preCfg, flags.HelpFlag)
	_, err := preParser.Parse()
	if err != nil {
		var flagsErr *flags.Error
		if ok := errors.As(err, &flagsErr); ok && flagsErr.Type == flags.ErrHelp {
			return nil, err
		}
	}

	appName := filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	usageMessage := fmt.Sprintf("Use %s -h to show usage", appName)

	// Show the version and exit if the version flag was specified.
	if preCfg.ShowVersion {
		fmt.Println(appName, "version", version.Version())
		os.Exit(0)
	}

	// Load additional config from file.
	var configFileError error
	parser := newConfigParser(cfgFlags, flags.Default)
	cfg := &Config{
		Flags: cfgFlags,
	}

	// Parse command line options again to ensure they take precedence.
	_, err = parser.Parse()
	if err != nil {
		var flagsErr *flags.Error
		if ok := errors.As(err, &flagsErr); !ok || flagsErr.Type != flags.ErrHelp {
			return nil, errors.Wrapf(err, "Error parsing command line arguments: %s\n\n%s", err, usageMessage)
		}
		return nil, err
	}

	// Create the home directory if it doesn't already exist.
	funcName := "loadConfig"
	err = os.MkdirAll(DefaultAppDir, 0700)
	if err != nil {
		// Show a nicer error message if it's because a symlink is
		// linked to a directory that does not exist (probably because
		// it's not mounted).
		var e *os.PathError
		if ok := errors.As(err, &e); ok && os.IsExist(err) {
			if link, lerr := os.Readlink(e.Path); lerr == nil {
				str := "is symlink %s -> %s mounted?"
				err = errors.Errorf(str, e.Path, link)
			}
		}

		str := "%s: Failed to create home directory: %s"
		err := errors.Errorf(str, funcName, err)
		return nil, err
	}

	err = cfg.ResolveNetwork(parser)
	if err != nil {
		return nil, err
	}

	cfg.AppDir = cleanAndExpandPath(cfg.AppDir)
	// Append the network type to the app directory so it is "namespaced"
	// per network.
	// All data is specific to a network, so namespacing the data directory
	// means each individual piece of serialized data does not have to
	// worry about changing names per network and such.
	cfg.AppDir = filepath.Join(cfg.AppDir, cfg.NetParams().Name)

	// Logs directory is usually under the home directory, unless otherwise specified
	if cfg.LogDir == "" {
		cfg.LogDir = filepath.Join(cfg.AppDir, defaultLogDirname)
	}
	cfg.LogDir = cleanAndExpandPath(cfg.LogDir)

	// Special show command to list supported subsystems and exit.
	if cfg.LogLevel == "show" {
		fmt.Println("Supported subsystems", logger.SupportedSubsystems())
		os.Exit(0)
	}

	// Initialize log rotation. After log rotation has been initialized, the
	// logger variables may be used.
	logger.InitLog(filepath.Join(cfg.LogDir, defaultLogFilename), filepath.Join(cfg.LogDir, defaultErrLogFilename))

	// Parse, validate, and set debug log level(s).
	if err := logger.ParseAndSetLogLevels(cfg.LogLevel); err != nil {
		err := errors.Errorf("%s: %s", funcName, err.Error())
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// Validate profile port number
	if cfg.Profile != "" {
		profilePort, err := strconv.Atoi(cfg.Profile)
		if err != nil || profilePort < 1024 || profilePort > 65535 {
			str := "%s: The profile port must be between 1024 and 65535"
			err := errors.Errorf(str, funcName)
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, usageMessage)
			return nil, err
		}
	}

	// Don't allow ban durations that are too short.
	if cfg.BanDuration < time.Second {
		str := "%s: The banduration option may not be less than 1s -- parsed [%s]"
		err := errors.Errorf(str, funcName, cfg.BanDuration)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// Validate any given whitelisted IP addresses and networks.
	if len(cfg.Whitelists) > 0 {
		var ip net.IP
		cfg.Whitelists = make([]*net.IPNet, 0, len(cfg.Flags.Whitelists))

		for _, addr := range cfg.Flags.Whitelists {
			_, ipnet, err := net.ParseCIDR(addr)
			if err != nil {
				ip = net.ParseIP(addr)
				if ip == nil {
					str := "%s: The whitelist value of '%s' is invalid"
					err = errors.Errorf(str, funcName, addr)
					fmt.Fprintln(os.Stderr, err)
					fmt.Fprintln(os.Stderr, usageMessage)
					return nil, err
				}
				var bits int
				if ip.To4() == nil {
					// IPv6
					bits = 128
				} else {
					bits = 32
				}
				ipnet = &net.IPNet{
					IP:   ip,
					Mask: net.CIDRMask(bits, bits),
				}
			}
			cfg.Whitelists = append(cfg.Whitelists, ipnet)
		}
	}

	// --addPeer and --connect do not mix.
	if len(cfg.AddPeers) > 0 && len(cfg.ConnectPeers) > 0 {
		str := "%s: the --addpeer and --connect options can not be " +
			"mixed"
		err := errors.Errorf(str, funcName)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// --proxy or --connect without --listen disables listening.
	if (cfg.Proxy != "" || len(cfg.ConnectPeers) > 0) &&
		len(cfg.Listeners) == 0 {
		cfg.DisableListen = true
	}

	// ConnectPeers means no DNS seeding and no outbound peers
	if len(cfg.ConnectPeers) > 0 {
		cfg.DisableDNSSeed = true
		cfg.TargetOutboundPeers = 0
	}

	// Add the default listener if none were specified. The default
	// listener is all addresses on the listen port for the network
	// we are to connect to.
	if len(cfg.Listeners) == 0 {
		cfg.Listeners = []string{
			net.JoinHostPort("", cfg.NetParams().DefaultPort),
		}
	}

	if cfg.DisableRPC {
		log.Infof("RPC service is disabled")
	}

	// Add the default RPC listener if none were specified. The default
	// RPC listener is all addresses on the RPC listen port for the
	// network we are to connect to.
	if !cfg.DisableRPC && len(cfg.RPCListeners) == 0 {
		cfg.RPCListeners = []string{
			net.JoinHostPort("", cfg.NetParams().RPCPort),
		}
	}

	if cfg.RPCMaxConcurrentReqs < 0 {
		str := "%s: The rpcmaxwebsocketconcurrentrequests option may " +
			"not be less than 0 -- parsed [%d]"
		err := errors.Errorf(str, funcName, cfg.RPCMaxConcurrentReqs)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// Validate the the minrelaytxfee.
	cfg.MinRelayTxFee, err = util.NewAmount(cfg.Flags.MinRelayTxFee)
	if err != nil {
		str := "%s: invalid minrelaytxfee: %s"
		err := errors.Errorf(str, funcName, err)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// Disallow 0 and negative min tx fees.
	if cfg.MinRelayTxFee == 0 {
		str := "%s: The minrelaytxfee option must be greater than 0 -- parsed [%d]"
		err := errors.Errorf(str, funcName, cfg.MinRelayTxFee)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// Look for illegal characters in the user agent comments.
	for _, uaComment := range cfg.UserAgentComments {
		if strings.ContainsAny(uaComment, "/:()") {
			err := errors.Errorf("%s: The following characters must not "+
				"appear in user agent comments: '/', ':', '(', ')'",
				funcName)
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, usageMessage)
			return nil, err
		}
	}

	// Add default port to all listener addresses if needed and remove
	// duplicate addresses.
	cfg.Listeners, err = network.NormalizeAddresses(cfg.Listeners,
		cfg.NetParams().DefaultPort)
	if err != nil {
		return nil, err
	}

	// Add default port to all rpc listener addresses if needed and remove
	// duplicate addresses.
	cfg.RPCListeners, err = network.NormalizeAddresses(cfg.RPCListeners,
		cfg.NetParams().RPCPort)
	if err != nil {
		return nil, err
	}

	// Disallow --addpeer and --connect used together
	if len(cfg.AddPeers) > 0 && len(cfg.ConnectPeers) > 0 {
		str := "%s: --addpeer and --connect can not be used together"
		err := errors.Errorf(str, funcName)
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usageMessage)
		return nil, err
	}

	// Add default port to all added peer addresses if needed and remove
	// duplicate addresses.
	cfg.AddPeers, err = network.NormalizeAddresses(cfg.AddPeers,
		cfg.NetParams().DefaultPort)
	if err != nil {
		return nil, err
	}

	cfg.ConnectPeers, err = network.NormalizeAddresses(cfg.ConnectPeers,
		cfg.NetParams().DefaultPort)
	if err != nil {
		return nil, err
	}

	// Setup dial and DNS resolution (lookup) functions depending on the
	// specified options. The default is to use the standard
	// net.DialTimeout function as well as the system DNS resolver. When a
	// proxy is specified, the dial function is set to the proxy specific
	// dial function.
	cfg.Dial = net.DialTimeout
	cfg.Lookup = net.LookupIP
	if cfg.Proxy != "" {
		_, _, err := net.SplitHostPort(cfg.Proxy)
		if err != nil {
			str := "%s: Proxy address '%s' is invalid: %s"
			err := errors.Errorf(str, funcName, cfg.Proxy, err)
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprintln(os.Stderr, usageMessage)
			return nil, err
		}

		proxy := &socks.Proxy{
			Addr:     cfg.Proxy,
			Username: cfg.ProxyUser,
			Password: cfg.ProxyPass,
		}
		cfg.Dial = proxy.DialTimeout
	}

	// Warn about missing config file only after all other configuration is
	// done. This prevents the warning on help messages and invalid
	// options. Note this should go directly before the return.
	if configFileError != nil {
		log.Warnf("%s", configFileError)
	}
	return cfg, nil
}

// Config defines the configuration options for Metchaind.
//
// See loadConfig for details on the configuration load process.
type Config struct {
	*Flags
	Lookup        func(string) ([]net.IP, error)
	Dial          func(string, string, time.Duration) (net.Conn, error)
	MiningAddrs   []util.Address
	MinRelayTxFee util.Amount
	Whitelists    []*net.IPNet
	SubnetworkID  *externalapi.DomainSubnetworkID // nil in full nodes
}

// cleanAndExpandPath expands environment variables and leading ~ in the
// passed path, cleans the result, and returns it.
func cleanAndExpandPath(path string) string {
	// Expand initial ~ to OS specific home directory.
	if strings.HasPrefix(path, "~") {
		homeDir := filepath.Dir(DefaultAppDir)
		path = strings.Replace(path, "~", homeDir, 1)
	}

	// NOTE: The os.ExpandEnv doesn't work with Windows-style %VARIABLE%,
	// but they variables can still be expanded via POSIX-style $VARIABLE.
	return filepath.Clean(os.ExpandEnv(path))
}
