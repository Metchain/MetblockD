package subnetworks

import (
	"github.com/Metchain/Metblock/external"
)

var (
	// SubnetworkIDNative is the default subnetwork ID which is used for transactions without related payload data
	SubnetworkIDNative = external.DomainSubnetworkID{}

	// SubnetworkIDCoinbase is the subnetwork ID which is used for the coinbase transaction
	SubnetworkIDCoinbase = external.DomainSubnetworkID{1}

	// SubnetworkIDRegistry is the subnetwork ID which is used for adding new sub networks to the registry
	SubnetworkIDRegistry = external.DomainSubnetworkID{2}
)
