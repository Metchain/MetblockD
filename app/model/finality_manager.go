package model

import "github.com/Metchain/MetblockD/external"

// FinalityManager provides method to validate that a block does not violate finality
type FinalityManager interface {
	VirtualFinalityPoint(stagingArea *StagingArea) (*external.DomainHash, error)
	FinalityPoint(stagingArea *StagingArea, blockHash *external.DomainHash, isBlockWithTrustedData bool) (*external.DomainHash, error)
}
