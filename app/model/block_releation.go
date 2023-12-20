package model

import "github.com/Metchain/MetblockD/external"

// BlockRelations represents a block's parent/child relations
type BlockRelations struct {
	Parents    []*external.DomainHash
	Children   []*external.DomainHash
	ParentMet  []*external.DomainHash
	ParentMega []*external.DomainHash
}
