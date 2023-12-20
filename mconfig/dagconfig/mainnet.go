package dagconfig

import (
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/mconfig"
)

// MainnetParams defines the network parameters for the main Metchain network.
var MainnetParams = Params{

	Name:        "metchain-mainnet",
	Net:         appmessage.Mainnet,
	RPCPort:     mconfig.MainnetPortRPC,
	DefaultPort: mconfig.MainnetPortP2p,
	DNSSeeds: []string{
		// This DNS seeder is run by GH
		"n.metminingpool.com",
		"n.metscan.io",
		"n.metchain.community",
		// This DNS seeder is run by GH

	},

	MetChainMiniBlock:  mconfig.MINING_REWARD,
	MetChainMiegaBlock: mconfig.MINING_REWARD_MEGA,
	MetChainMetBlock:   mconfig.MINING_REWARD_MET,
	DeadWallet:         mconfig.DeadWallet,

	TargetTimePerBlock: mconfig.MINING_TIMER_SEC,
}
