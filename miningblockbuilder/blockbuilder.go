package miningblockbuilder

import (
	"github.com/Metchain/Metblock/blockchain/consensus/consensusreference"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/miningblockbuilder/miningmempoolmodel"
	"github.com/Metchain/Metblock/utils/difficulty"
	"github.com/Metchain/Metblock/utils/logger"
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
