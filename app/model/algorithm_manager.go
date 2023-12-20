package model

import "github.com/Metchain/MetblockD/external"

// DifficultyManager provides a method to resolve the
// difficulty value of a block
type AlgorithmManager interface {
	GetActiveMiningAlgo(stagingArea *StagingArea, blockHash *external.DomainHash, isBlockWithTrustedData bool) (uint32, error)
	VerifyMiningAlgo(stagingArea *StagingArea, blockHash *external.DomainHash) (uint32, error)
	BroadcastMiningAlgo(startHash *external.DomainHash, windowSize int) (uint64, error)
	SetAlgoBlockId()
	VerifyAlgoBlockId()
	SetDifficultyRetargetForAlgo()
	GetDifficultyRetargetForAlgo()
	VerifyDifficultyRetargetForAlgo()
	BroadcastDifficultyRetargetForAlgo()
	ConcensusAntiAsics()
}
