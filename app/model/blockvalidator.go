package model

import "github.com/Metchain/Metblock/external"

// BlockValidator exposes a set of validation classes, after which
// it's possible to determine whether a block is valid
type BlockValidator interface {
	ValidateHeaderInIsolation(stagingArea *StagingArea, blockHash *external.DomainHash) error
	ValidateBodyInIsolation(stagingArea *StagingArea, blockHash *external.DomainHash) error
	ValidateHeaderInContext(stagingArea *StagingArea, blockHash *external.DomainHash, isBlockWithTrustedData bool) error
	ValidateBodyInContext(stagingArea *StagingArea, blockHash *external.DomainHash, isBlockWithTrustedData bool) error
	ValidatePruningPointViolationAndProofOfWorkAndDifficulty(stagingArea *StagingArea, blockHash *external.DomainHash, isBlockWithTrustedData bool) error
}
