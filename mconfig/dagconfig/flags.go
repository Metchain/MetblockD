package dagconfig

import (
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/mconfig"
	"time"
)

type Params struct {
	M external.MType

	// Name defines a human-readable identifier for the network.
	Name string

	// Net defines the magic bytes used to identify the network.
	Net appmessage.MetchainNet

	// RPCPort defines the rpc server port
	RPCPort string

	// DefaultPort defines the default peer-to-peer port for the network.
	DefaultPort string

	// DNSSeeds defines a list of DNS seeds for the network that are used
	// as one method to discover peers.
	DNSSeeds []string

	// GRPCSeeds defines a list of GRPC seeds for the network that are used
	// as one method to discover peers.
	GRPCSeeds []string

	// TargetTimePerBlock is the desired amount of time to generate each
	// block.
	TargetTimePerBlock time.Duration

	MetChainMiniBlock  float64
	MetChainMiegaBlock float64
	MetChainMetBlock   float64
	DeadWallet         string

	Prefix mconfig.Bech32Prefix
}
