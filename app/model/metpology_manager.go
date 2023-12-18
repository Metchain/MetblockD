package model

import "github.com/Metchain/Metblock/external"

type MetTopologyManager interface {
	Parents(stagingArea *StagingArea, blockHash *external.DomainHash) ([]*external.DomainHash, error)
	Children(stagingArea *StagingArea, blockHash *external.DomainHash) ([]*external.DomainHash, error)
	IsParentOf(stagingArea *StagingArea, blockHashA *external.DomainHash, blockHashB *external.DomainHash) (bool, error)
	IsChildOf(stagingArea *StagingArea, blockHashA *external.DomainHash, blockHashB *external.DomainHash) (bool, error)
	IsAncestorOf(stagingArea *StagingArea, blockHashA *external.DomainHash, blockHashB *external.DomainHash) (bool, error)
	IsAncestorOfAny(stagingArea *StagingArea, blockHash *external.DomainHash, potentialDescendants []*external.DomainHash) (bool, error)
	IsAnyAncestorOf(stagingArea *StagingArea, potentialAncestors []*external.DomainHash, blockHash *external.DomainHash) (bool, error)
	IsInSelectedParentChainOf(stagingArea *StagingArea, blockHashA *external.DomainHash, blockHashB *external.DomainHash) (bool, error)
	ChildInSelectedParentChainOf(stagingArea *StagingArea, lowHash, highHash *external.DomainHash) (*external.DomainHash, error)

	SetParents(stagingArea *StagingArea, blockHash *external.DomainHash, parentHashes []*external.DomainHash) error
}
