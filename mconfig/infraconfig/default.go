package infraconfig

import (
	_ "embed"
	"github.com/Metchain/MetblockD/mconfig/util"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultConfigFilename      = "metchaind.conf"
	defaultLogLevel            = "info"
	defaultLogDirname          = "logs"
	defaultLogFilename         = "metchaind.log"
	defaultErrLogFilename      = "metchaind_err.log"
	defaultDataDirname         = "datadir"
	defaultTargetOutboundPeers = 8
	defaultMaxInboundPeers     = 117
	defaultBanDuration         = time.Hour * 24
	defaultBanThreshold        = 100
	//DefaultConnectTimeout is the default connection timeout when dialing
	DefaultConnectTimeout = time.Second * 30
	//DefaultMaxRPCClients is the default max number of RPC clients
	DefaultMaxRPCClients         = 128
	defaultMaxRPCWebsockets      = 25
	defaultMaxRPCConcurrentReqs  = 20
	defaultBlockMaxMass          = 10_000_000
	blockMaxMassMin              = 1000
	blockMaxMassMax              = 10_000_000
	defaultMinRelayTxFee         = 1e-5 // 1 sompi per byte
	defaultMaxOrphanTransactions = 100
	//DefaultMaxOrphanTxSize is the default maximum size for an orphan transaction
	DefaultMaxOrphanTxSize  = 100_000
	defaultSigCacheMaxSize  = 100_000
	sampleConfigFilename    = "sample-metchain.conf"
	defaultMaxUTXOCacheSize = 5_000_000_000
	defaultProtocolVersion  = 5
)

var (
	// DefaultAppDir is the default home directory for metchaind.
	DefaultAppDir = util.AppDir("metchaind", false)

	defaultConfigFile = filepath.Join(DefaultAppDir, defaultConfigFilename)
	defaultDataDir    = filepath.Join(DefaultAppDir)

	defaultRPCKeyFile  = filepath.Join(DefaultAppDir, "rpc.key")
	defaultRPCCertFile = filepath.Join(DefaultAppDir, "rpc.cert")
)

func defaultFlags() *Flags {
	return &Flags{
		ConfigFile:           defaultConfigFile,
		LogLevel:             defaultLogLevel,
		DataDir:              defaultDataDirname,
		TargetOutboundPeers:  defaultTargetOutboundPeers,
		MaxInboundPeers:      defaultMaxInboundPeers,
		BanDuration:          defaultBanDuration,
		BanThreshold:         defaultBanThreshold,
		RPCMaxClients:        DefaultMaxRPCClients,
		RPCMaxWebsockets:     defaultMaxRPCWebsockets,
		RPCMaxConcurrentReqs: defaultMaxRPCConcurrentReqs,
		AppDir:               defaultDataDir,
		RPCKey:               defaultRPCKeyFile,
		RPCCert:              defaultRPCCertFile,
		BlockMaxMass:         defaultBlockMaxMass,
		MaxOrphanTxs:         defaultMaxOrphanTransactions,
		SigCacheMaxSize:      defaultSigCacheMaxSize,
		MinRelayTxFee:        defaultMinRelayTxFee,
		MaxUTXOCacheSize:     defaultMaxUTXOCacheSize,
		ServiceOptions:       &ServiceOptions{},
		ProtocolVersion:      defaultProtocolVersion,
	}
}

// createDefaultConfig copies the file sample-metchain.conf to the given destination path,
// and populates it with some randomly generated RPC username and password.
func createDefaultConfigFile(destinationPath string) error {
	// Create the destination directory if it does not exists
	err := os.MkdirAll(filepath.Dir(destinationPath), 0700)
	if err != nil {
		return err
	}

	dest, err := os.OpenFile(destinationPath,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = dest.WriteString(sampleConfig)

	return err
}

//go:embed sample-metchain.conf
var sampleConfig string
