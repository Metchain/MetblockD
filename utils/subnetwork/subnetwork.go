package subnetworks

import (
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/mconfig/externalapi"
	"github.com/pkg/errors"
)

// FromBytes creates a DomainSubnetworkID from the given byte slice
func FromBytes(subnetworkIDBytes []byte) (*external.DomainSubnetworkID, error) {
	if len(subnetworkIDBytes) != externalapi.DomainSubnetworkIDSize {
		return nil, errors.Errorf("invalid hash size. Want: %d, got: %d",
			external.DomainSubnetworkIDSize, len(subnetworkIDBytes))
	}
	var domainSubnetworkID external.DomainSubnetworkID
	copy(domainSubnetworkID[:], subnetworkIDBytes)
	return &domainSubnetworkID, nil
}
