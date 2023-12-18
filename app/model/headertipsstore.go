package model

import "github.com/Metchain/Metblock/external"

// HeaderSelectedTipStore represents a store of the headers selected tip
type HeaderSelectedTipStore interface {
	Stage(stagingArea *StagingArea, selectedTip *external.DomainHash)
	IsStaged(stagingArea *StagingArea) bool
	HeadersSelectedTip(dbContext DBReader, stagingArea *StagingArea) (*external.DomainHash, error)
	Has(dbContext DBReader, stagingArea *StagingArea) (bool, error)
}
