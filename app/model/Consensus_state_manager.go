package model

import "github.com/Metchain/MetblockD/external"

type ConsensusStateManager interface {
	// Add Block function to be created
	AddBlock(stagingArea *StagingArea, blockHash *external.DomainHash, updateVirtual bool) (*external.SelectedChainPath, external.UTXODiff, *UTXODiffReversalData, error)
	PopulateTransactionWithUTXOEntries(stagingArea *StagingArea, transaction *external.DomainTransaction) error

	RestorePastUTXOSetIterator(stagingArea *StagingArea, blockHash *external.DomainHash) (external.ReadOnlyUTXOSetIterator, error)

	GetVirtualSelectedParentChainFromBlock(stagingArea *StagingArea, blockHash *external.DomainHash) (*external.SelectedChainPath, error)
}
