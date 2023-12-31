package model

import "github.com/Metchain/MetblockD/external"

// HeadersSelectedTipManager manages the state of the headers selected tip
type HeadersSelectedTipManager interface {
	AddHeaderTip(stagingArea *StagingArea, hash *external.DomainHash) error
}
