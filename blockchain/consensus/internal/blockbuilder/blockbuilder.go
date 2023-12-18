package blockbuilder

import (
	"github.com/Metchain/Metblock/blockchain/consensus/blockheader"
	"github.com/Metchain/Metblock/blockchain/consensus/stagingarea"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/utils/logger"
)

type blockBuilder struct {
	databaseContext database.Database
}

func New(db database.Database) *blockBuilder {
	return &blockBuilder{
		databaseContext: db,
	}
}

func (bb *blockBuilder) BuildBlockTemplate(coinbaseData *external.DomainCoinbaseData,
	transactions []*external.DomainTransaction) (block *external.DomainBlock, coinbaseHasRedReward bool, err error) {

	onEnd := logger.LogAndMeasureExecutionTime(log, "BuildBlock")
	defer onEnd()

	return bb.buildBlocktemplate(coinbaseData, transactions)
}

func (bb *blockBuilder) buildBlocktemplate(coinbaseData *external.DomainCoinbaseData,
	transactions []*external.DomainTransaction) (block *external.DomainBlock, coinbaseHasRedReward bool, err error) {

	/*err = bb.validateTransactions(stagingArea, transactions)
	if err != nil {
		return nil, false, err
	}

	coinbase, coinbaseHasRedReward, err := bb.newBlockCoinbaseTransaction(stagingArea, coinbaseData)
	if err != nil {
		return nil, false, err
	}
	transactionsWithCoinbase := append([]*external.DomainTransaction{coinbase}, transactions...)*/
	transactionsWithCoinbase := []*external.DomainTransaction{}

	header, err := bb.buildHeaderTemplate(coinbaseData, transactionsWithCoinbase)
	if err != nil {
		return nil, false, err
	}

	return &external.DomainBlock{
		Header:       header,
		Transactions: transactionsWithCoinbase,
	}, coinbaseHasRedReward, nil
}

func (bb *blockBuilder) buildHeaderTemplate(coinbaseData *external.DomainCoinbaseData, transactions []*external.DomainTransaction) (external.BlockHeader, error) {

	staging := stagingarea.NewStagingArea(bb.databaseContext)
	stagingBlock := staging.StagingBlock()
	Parents := []*external.BlockLevelParents{
		&external.BlockLevelParents{
			stagingBlock.Merkleroot,
		},
		&external.BlockLevelParents{
			stagingBlock.Blockhash,
		},
		&external.BlockLevelParents{
			stagingBlock.Metblock,
		},
		&external.BlockLevelParents{
			stagingBlock.Megablock,
		},
	}

	utxoCommitment := bb.getBlockUTXOCommitment(coinbaseData)

	return blockheader.NewImmutableBlockHeader(
		0,
		stagingBlock.Blockheight,
		stagingBlock.Blockhash,
		stagingBlock.Previoushash,
		stagingBlock.Merkleroot,
		Parents,
		stagingBlock.Metblock,
		stagingBlock.Megablock,
		stagingBlock.Childblocks,
		stagingBlock.TimeInMilliseconds,
		stagingBlock.Bits,
		stagingBlock.Nonce,
		stagingBlock.BlockLevel,
		stagingBlock.VerificationPoint,
		utxoCommitment,
	), nil

}

func (bb *blockBuilder) getBlockUTXOCommitment(coinbaseData *external.DomainCoinbaseData) []byte {
	return coinbaseData.ScriptPublicKey.Script
}
