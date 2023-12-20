package infraconfig

import (
	"github.com/Metchain/MetblockD/mconfig/dagconfig"
	"time"
)

type Flags struct {
	ShowVersion                     bool          `short:"V" long:"version" description:"Display version information and exit"`
	ConfigFile                      string        `short:"C" long:"configfile" description:"Path to configuration file"`
	AppDir                          string        `short:"b" long:"appdir" description:"Directory to store data"`
	DataDir                         string        `short:"D" long:"datadir" description:"Directory to store info"`
	LogDir                          string        `long:"logdir" description:"Directory to log output."`
	AddPeers                        []string      `short:"a" long:"addpeer" description:"Add a peer to connect with at startup"`
	ConnectPeers                    []string      `long:"connect" description:"Connect only to the specified peers at startup"`
	DisableListen                   bool          `long:"nolisten" description:"Disable listening for incoming connections -- NOTE: Listening is automatically disabled if the --connect or --proxy options are used without also specifying listen interfaces via --listen"`
	Listeners                       []string      `long:"listen" description:"Add an interface/port to listen for connections (default all interfaces port: 16111, testnet: 16211)"`
	TargetOutboundPeers             int           `long:"outpeers" description:"Target number of outbound peers"`
	MaxInboundPeers                 int           `long:"maxinpeers" description:"Max number of inbound peers"`
	EnableBanning                   bool          `long:"enablebanning" description:"Enable banning of misbehaving peers"`
	BanDuration                     time.Duration `long:"banduration" description:"How long to ban misbehaving peers. Valid time units are {s, m, h}. Minimum 1 second"`
	BanThreshold                    uint32        `long:"banthreshold" description:"Maximum allowed ban score before disconnecting and banning misbehaving peers."`
	Whitelists                      []string      `long:"whitelist" description:"Add an IP network or IP that will not be banned. (eg. 192.168.1.0/24 or ::1)"`
	RPCListeners                    []string      `long:"rpclisten" description:"Add an interface/port to listen for RPC connections (default port: 16110, testnet: 16210)"`
	RPCCert                         string        `long:"rpccert" description:"File containing the certificate file"`
	RPCKey                          string        `long:"rpckey" description:"File containing the certificate key"`
	RPCMaxClients                   int           `long:"rpcmaxclients" description:"Max number of RPC clients for standard connections"`
	RPCMaxWebsockets                int           `long:"rpcmaxwebsockets" description:"Max number of RPC websocket connections"`
	RPCMaxConcurrentReqs            int           `long:"rpcmaxconcurrentreqs" description:"Max number of concurrent RPC requests that may be processed concurrently"`
	DisableRPC                      bool          `long:"norpc" description:"Disable built-in RPC server"`
	SafeRPC                         bool          `long:"saferpc" description:"Disable RPC commands which affect the state of the node"`
	DisableDNSSeed                  bool          `long:"nodnsseed" description:"Disable DNS seeding for peers"`
	DNSSeed                         string        `long:"dnsseed" description:"Override DNS seeds with specified hostname (Only 1 hostname allowed)"`
	GRPCSeed                        string        `long:"grpcseed" description:"Hostname of gRPC server for seeding peers"`
	ExternalIPs                     []string      `long:"externalip" description:"Add an ip to the list of local addresses we claim to listen on to peers"`
	Proxy                           string        `long:"proxy" description:"Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)"`
	ProxyUser                       string        `long:"proxyuser" description:"Username for proxy server"`
	ProxyPass                       string        `long:"proxypass" default-mask:"-" description:"Password for proxy server"`
	DbType                          string        `long:"dbtype" description:"Database backend to use for the Block DAG"`
	Profile                         string        `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
	LogLevel                        string        `short:"d" long:"loglevel" description:"Logging level for all subsystems {trace, debug, info, warn, error, critical} -- You may also specify <subsystem>=<level>,<subsystem2>=<level>,... to set the log level for individual subsystems -- Use show to list available subsystems"`
	Upnp                            bool          `long:"upnp" description:"Use UPnP to map our listening port outside of NAT"`
	MinRelayTxFee                   float64       `long:"minrelaytxfee" description:"The minimum transaction fee in Met to be considered a non-zero fee."`
	MaxOrphanTxs                    uint64        `long:"maxorphantx" description:"Max number of orphan transactions to keep in memory"`
	BlockMaxMass                    uint64        `long:"blockmaxmass" description:"Maximum transaction mass to be used when creating a block"`
	UserAgentComments               []string      `long:"uacomment" description:"Comment to add to the user agent -- See BIP 14 for more information."`
	NoPeerBloomFilters              bool          `long:"nopeerbloomfilters" description:"Disable bloom filtering support"`
	SigCacheMaxSize                 uint          `long:"sigcachemaxsize" description:"The maximum number of entries in the signature verification cache"`
	BlocksOnly                      bool          `long:"blocksonly" description:"Do not accept transactions from remote peers."`
	RelayNonStd                     bool          `long:"relaynonstd" description:"Relay non-standard transactions regardless of the default settings for the active network."`
	RejectNonStd                    bool          `long:"rejectnonstd" description:"Reject non-standard transactions regardless of the default settings for the active network."`
	ResetDatabase                   bool          `long:"reset-db" description:"Reset database before starting node. It's needed when switching between subnetworks."`
	MaxUTXOCacheSize                uint64        `long:"maxutxocachesize" description:"Max size of loaded UTXO into ram from the disk in bytes"`
	UTXOIndex                       bool          `long:"utxoindex" description:"Enable the UTXO index"`
	IsArchivalNode                  bool          `long:"archival" description:"Run as an archival node: don't delete old block data when moving the pruning point (Warning: heavy disk usage)'"`
	AllowSubmitBlockWhenNotSynced   bool          `long:"allow-submit-block-when-not-synced" hidden:"true" description:"Allow the node to accept blocks from RPC while not synced (this flag is mainly used for testing)"`
	EnableSanityCheckPruningUTXOSet bool          `long:"enable-sanity-check-pruning-utxo" hidden:"true" description:"When moving the pruning point - check that the utxo set matches the utxo commitment"`
	ProtocolVersion                 uint32        `long:"protocol-version" description:"Use non default p2p protocol version"`
	NetworkFlags
	ServiceOptions *ServiceOptions
}

type NetworkFlags struct {
	Testnet         bool `long:"testnet" description:"Use the test network"`
	ActiveNetParams *dagconfig.Params
}
