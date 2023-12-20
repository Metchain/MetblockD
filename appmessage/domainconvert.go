package appmessage

import (
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/utils/hashes"
)

func DomainBlockToRPCBlock(block *external.DomainBlock) (*RPCBlock, *external.DomainBlock) {
	// update this now Metchainupdate

	parents := make([]*RPCBlockLevelParents, len(block.Header.Parents()))
	for i, blockLevelParents := range block.Header.Parents() {
		parents[i] = &RPCBlockLevelParents{
			ParentHashes: hashes.ToStrings(blockLevelParents),
		}
	}

	header := &RPCBlockHeader{
		Version:              uint32(block.Header.Version()),
		Parents:              parents,
		HashMerkleRoot:       block.Header.Previoushash().String(),
		AcceptedIDMerkleRoot: block.Header.Previoushash().String(),
		UTXOCommitment:       string(block.Header.UtxoCommitment()),
		Timestamp:            block.Header.TimeInMilliseconds(),
		Bits:                 uint32(block.Header.Bits()),
		Nonce:                block.Header.Nonce(),
	}

	return &RPCBlock{
		Header: header,
	}, block
}
