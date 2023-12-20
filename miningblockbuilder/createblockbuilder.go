package miningblockbuilder

import (
	"github.com/Metchain/MetblockD/blockchain/consensus/consensusreference"
	"github.com/Metchain/MetblockD/blockchain/domainconsensus/miningmanager/miningmodel"
)

// New creates a new blockTemplateBuilder
func New(consensusReference consensusreference.ConsensusReference, mempool miningmodel.Mempool) miningmodel.BlockTemplateBuilder {
	return &blockTemplateBuilder{
		consensusReference: consensusReference,
		mempool:            mempool,
	}
}
