package consensus

import (
	"github.com/Metchain/MetblockD/app/model"
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/staging"
	"sync"
)

type consensus struct {
	lock            *sync.Mutex
	databaseContext database.Database

	//Update domain block
	genesisBlock *external.DomainBlock
	genesisHash  *external.DomainHash

	expectedDAAWindowDurationInMilliseconds int64

	//Update Block Processor
	blockProcessor model.BlockProcessor
	//Update Block Builder
	blockBuilder          model.BlockBuilder
	consensusStateManager model.ConsensusStateManager
	transactionValidator  model.TransactionValidator
	syncManager           model.SyncManager
	pastMedianTimeManager model.PastMedianTimeManager
	blockValidator        model.BlockValidator
	coinbaseManager       model.CoinbaseManager

	difficultyManager model.DifficultyManager

	//For Advance Mining ALgorithm
	algorithmManager model.AlgorithmManager

	headerTipsManager model.HeadersSelectedTipManager

	finalityManager model.FinalityManager

	blockStore model.BlockStore

	blockRelationStores []model.BlockRelationStore
	blockStatusStore    model.BlockStatusStore

	headersSelectedTipStore model.HeaderSelectedTipStore
	multisetStore           model.MultisetStore

	finalityStore             model.FinalityStore
	headersSelectedChainStore model.HeadersSelectedChainStore

	consensusEventsChan chan external.ConsensusEvent
	virtualNotUpdated   bool
}

func (c *consensus) BuildBlock(block *external.TempBlock) (*external.TempBlock, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	block, err := c.blockBuilder.BuildBlock(block)

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (s *consensus) ValidateAndInsertBlock(block *external.DomainBlock, updateVirtual bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) ValidateTransactionAndPopulateWithConsensusData(transaction *external.DomainTransaction) error {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) GetBlock(blockHash *external.DomainHash) (*external.DomainBlock, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) GetBlockEvenIfHeaderOnly(blockHash *external.DomainHash) (*external.DomainBlock, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) GetBlockHeader(blockHash *external.DomainHash) (external.BlockHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) GetBlockInfo(blockHash *external.DomainHash) (*external.BlockInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) GetBlockRelations(blockHash *external.DomainHash) (parents []*external.DomainHash, children []*external.DomainHash, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) GetSyncInfo() (*external.SyncInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) IsValidVerificationPoint(blockHash *external.DomainHash) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) IsValidBlockchainUpdatePoint(blockHash *external.DomainHash) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) IsInSelectedParentChainOf(blockHashA *external.DomainHash, blockHashB *external.DomainHash) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) EstimateNetworkHashesPerSecond(startHash *external.DomainHash, windowSize int) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) IsChainBlock(blockHash *external.DomainHash) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) IsNearlySynced() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *consensus) Init(skipAddingGenesis bool) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	onEnd := logger.LogAndMeasureExecutionTime(log, "Init")
	defer onEnd()

	stagingArea := model.NewStagingArea()

	//Add Genesis verification here.
	exists, err := s.blockStatusStore.Exists(s.databaseContext, stagingArea, model.VirtualGenesisBlockHash)
	if err != nil {
		return err
	}

	if !exists {
		s.blockStatusStore.Stage(stagingArea, model.VirtualGenesisBlockHash, external.StatusUTXOValid)
		err = staging.CommitAllChanges(s.databaseContext, stagingArea)
		if err != nil {
			return err
		}
	}

	if !skipAddingGenesis && s.blockStore.Count(stagingArea) == 0 {
		genesisWithTrustedData := &external.BlockWithTrustedData{
			Block: s.genesisBlock,

			MetDagData: []*external.BlockMetDataHashPair{
				{
					MetGDagData: external.NewMETGBlockData(model.VirtualGenesisBlockHash),
					Hash:        s.genesisHash,
				},
			},
		}
		_, _, err = s.blockProcessor.ValidateAndInsertBlockWithTrustedData(genesisWithTrustedData, true)
		if err != nil {
			return err
		}
	}

	return nil
}
func NewFactory() Factory {
	return new(factory)
}
