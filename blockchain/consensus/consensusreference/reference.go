package consensusreference

import "github.com/Metchain/Metblock/external"

// ConsensusReference holds a reference to a consensus object.
// The consensus object may be swapped with a new one entirely
// during the IBD process. Before an atomic consensus operation,
// callers are expected to call Consensus() once and work against
// that instance throughout.
type ConsensusReference struct {
	consensus **external.Consensus
}

// Consensus returns the underlying consensus
func (ref ConsensusReference) Consensus() external.Consensus {
	return **ref.consensus
}

// NewConsensusReference constructs a new ConsensusReference
func NewConsensusReference(consensus **external.Consensus) ConsensusReference {
	return ConsensusReference{consensus: consensus}
}
