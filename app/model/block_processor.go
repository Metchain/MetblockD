package model

import "github.com/Metchain/MetblockD/external"

type BlockProcessor interface {
	ValidateAndInsertBlock(block *external.DomainBlock, shouldValidateAgainstUTXO bool) (*external.VirtualChangeSet, external.BlockStatus, error)

	ValidateAndInsertBlockWithTrustedData(block *external.BlockWithTrustedData, validateUTXO bool) (*external.VirtualChangeSet, external.BlockStatus, error)
}
