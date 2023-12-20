package blockbuilder

import (
	"github.com/Metchain/MetblockD/blockchain/consensus/blockheader"
	"github.com/Metchain/MetblockD/blockchain/consensus/stagingarea"
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/utils/logger"
)

type blockBuilder struct {
	databaseContext database.Database
}

func New(db database.Database) *blockBuilder {
	return &blockBuilder{
		databaseContext: db,
	}
}

func (bb *blockBuilder) BuildBlock(tempblock *external.TempBlock) (block *external.TempBlock, err error) {

	onEnd := logger.LogAndMeasureExecutionTime(log, "BuildBlock")
	defer onEnd()
	staging := stagingarea.NewStagingArea(bb.databaseContext)
	dblock, err := staging.AddBlock(tempblock)
	if err != nil {
		return nil, err
	}

	return dblock, nil
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
	Parents := []external.BlockLevelParents{

		[]*external.DomainHash{stagingBlock.Metblock, stagingBlock.Megablock, stagingBlock.Blockhash},
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
