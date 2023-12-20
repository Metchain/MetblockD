package miningblockbuilder

import (
	"github.com/Metchain/MetblockD/blockchain/consensus/consensusreference"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/miningblockbuilder/miningmempoolmodel"
	"github.com/Metchain/MetblockD/utils/difficulty"
	"github.com/Metchain/MetblockD/utils/logger"
)

type blockTemplateBuilder struct {
	consensusReference consensusreference.ConsensusReference
	mempool            miningmempoolmodel.Mempool

	coinbasePayloadScriptPublicKeyMaxLength uint8
}

var log = logger.RegisterSubSystem("MetD-Miningblockbuilder")

func (btb *blockTemplateBuilder) BuildBlockTemplate(coinbaseData *external.DomainCoinbaseData) (*external.DomainBlockTemplate, error) {
	block, err := btb.consensusReference.Consensus().BuildBlockTemplate(coinbaseData)
	if err != nil {
		return nil, err
	}
	log.Infof("Created new block template (target difficulty %064x)", difficulty.CompactToBig(block.Block.Header.Bits()))
	return block, nil
}

func (btb *blockTemplateBuilder) ModifyBlockTemplate(newCoinbaseData *external.DomainCoinbaseData) (*external.DomainBlockTemplate, error) {
	//TODO implement me
	panic("implement me")
}

func (btb *blockTemplateBuilder) BuildBlock(block *external.TempBlock) (*external.TempBlock, error) {

	block, err := btb.consensusReference.Consensus().BuildBlock(block)
	if err != nil {
		return nil, err
	}

	return block, nil
}
