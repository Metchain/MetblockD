package external

import "math/big"

// BlockInfo contains various information about a specific block
type BlockInfo struct {
	Exists         bool
	BlockStatus    BlockStatus
	BlueScore      uint64
	BlueWork       *big.Int
	SelectedParent *DomainHash
	MergeSetBlues  []*DomainHash
	MergeSetReds   []*DomainHash
}
