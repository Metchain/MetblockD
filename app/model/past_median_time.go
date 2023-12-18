package model

import "github.com/Metchain/Metblock/external"

// PastMedianTimeManager provides a method to resolve the
// past median time of a block
type PastMedianTimeManager interface {
	PastMedianTime(stagingArea *StagingArea, blockHash *external.DomainHash) (int64, error)
	InvalidateVirtualPastMedianTimeCache()
}
