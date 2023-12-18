package model

import "github.com/Metchain/Metblock/external"

// UTXODiffReversalData is used by ConsensusStateManager to reverse the UTXODiffs during a re-org
type UTXODiffReversalData struct {
	SelectedParentHash     *external.DomainHash
	SelectedParentUTXODiff external.UTXODiff
}
