package model

import "github.com/Metchain/Metblock/external"

// DifficultyManager provides a method to resolve the
// difficulty value of a block
type DifficultyManager interface {
	StageDAADataAndReturnRequiredDifficulty(stagingArea *StagingArea, blockHash *external.DomainHash, isBlockWithTrustedData bool) (uint32, error)
	RequiredDifficulty(stagingArea *StagingArea, blockHash *external.DomainHash) (uint32, error)
	EstimateNetworkHashesPerSecond(startHash *external.DomainHash, windowSize int) (uint64, error)
}
