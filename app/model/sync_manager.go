package model

import "github.com/Metchain/Metblock/external"

type SyncManager interface {
	GetHashesBetween(stagingArea *StagingArea, lowHash, highHash *external.DomainHash, maxBlocks uint64) (
		hashes []*external.DomainHash, actualHighHash *external.DomainHash, err error)
	GetAnticone(stagingArea *StagingArea, blockHash, contextHash *external.DomainHash, maxBlocks uint64) (hashes []*external.DomainHash, err error)
	GetMissingBlockBodyHashes(stagingArea *StagingArea, highHash *external.DomainHash) ([]*external.DomainHash, error)
	CreateBlockLocator(stagingArea *StagingArea, lowHash, highHash *external.DomainHash, limit uint32) (
		external.BlockLocator, error)
	CreateHeadersSelectedChainBlockLocator(stagingArea *StagingArea, lowHash, highHash *external.DomainHash) (
		external.BlockLocator, error)
	GetSyncInfo(stagingArea *StagingArea) (*external.SyncInfo, error)
}
