package consensushashing

import (
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/utils/hashes"
	"github.com/Metchain/Metblock/utils/serialization"
	"io"

	"github.com/pkg/errors"
)

// BlockHash returns the given block's hash
func BlockHash(block *external.DomainBlock) *external.DomainHash {
	return HeaderHash(block.Header)
}

// HeaderHash returns the given header's hash
func HeaderHash(header external.BaseBlockHeader) *external.DomainHash {
	// Encode the header and hash everything prior to the number of
	// transactions.
	writer := hashes.NewBlockHashWriter()
	err := serializeHeader(writer, header)
	if err != nil {
		// It seems like this could only happen if the writer returned an error.
		// and this writer should never return an error (no allocations or possible failures)
		// the only non-writer error path here is unknown types in `WriteElement`
		panic(errors.Wrap(err, "this should never happen. Hash digest should never return an error"))
	}

	return writer.Finalize()
}

func serializeHeader(w io.Writer, header external.BaseBlockHeader) error {
	timestamp := header.TimeInMilliseconds()
	// Metchainupdate needed here.
	/*numParents := len(header.Parents())
	if err := serialization.WriteElements(w, header.Version(), uint64(numParents)); err != nil {
		return err
	}
	for _, blockLevelParents := range header.Parents() {
		numBlockLevelParents := len(blockLevelParents)
		if err := serialization.WriteElements(w, uint64(numBlockLevelParents)); err != nil {
			return err
		}
		for _, hash := range blockLevelParents {
			if err := serialization.WriteElement(w, hash); err != nil {
				return err
			}
		}
	}*/
	return serialization.WriteElements(w, header.Merkleroot(), header.Blockheight(), header.UTXOCommitment(), timestamp,
		header.Bits(), header.Nonce())
}
