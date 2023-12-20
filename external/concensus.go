package external

// ConsensusEvent is an interface type that is implemented by all events raised by consensus
type ConsensusEvent interface {
	isConsensusEvent()
}

type Consensus interface {
	Init(skipAddingGenesis bool) error
	//Needed for Mining
	BuildBlockTemplate(coinbaseData *DomainCoinbaseData) (*DomainBlockTemplate, error)
	BuildBlock(block *TempBlock) (*TempBlock, error)

	/*BuildBlock(coinbaseData *DomainCoinbaseData, transactions []*DomainTransaction) (*DomainBlock, error)
	ValidateAndInsertBlock(block *DomainBlock, updateVirtual bool) error
	ValidateTransactionAndPopulateWithConsensusData(transaction *DomainTransaction) error

	GetBlock(blockHash *DomainHash) (*DomainBlock, bool, error)
	GetBlockEvenIfHeaderOnly(blockHash *DomainHash) (*DomainBlock, error)
	GetBlockHeader(blockHash *DomainHash) (BlockHeader, error)
	GetBlockInfo(blockHash *DomainHash) (*BlockInfo, error)
	GetBlockRelations(blockHash *DomainHash) (parents []*DomainHash, children []*DomainHash, err error)
	GetSyncInfo() (*SyncInfo, error)
	IsInSelectedParentChainOf(blockHashA *DomainHash, blockHashB *DomainHash) (bool, error)
	EstimateNetworkHashesPerSecond(startHash *DomainHash, windowSize int) (uint64, error)
	IsChainBlock(blockHash *DomainHash) (bool, error)
	IsNearlySynced() (bool, error)*/
}
