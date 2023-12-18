package serialization

import (
	"github.com/Metchain/Metblock/external"
	subnetworks "github.com/Metchain/Metblock/utils/subnetwork"
)

// DbSubnetworkIDToDomainSubnetworkID converts DbSubnetworkId to DomainSubnetworkID
func DbSubnetworkIDToDomainSubnetworkID(dbSubnetworkID *DbSubnetworkId) (*external.DomainSubnetworkID, error) {
	return subnetworks.FromBytes(dbSubnetworkID.SubnetworkId)
}

// DomainSubnetworkIDToDbSubnetworkID converts DomainSubnetworkID to DbSubnetworkId
func DomainSubnetworkIDToDbSubnetworkID(domainSubnetworkID *external.DomainSubnetworkID) *DbSubnetworkId {
	return &DbSubnetworkId{SubnetworkId: domainSubnetworkID[:]}
}
