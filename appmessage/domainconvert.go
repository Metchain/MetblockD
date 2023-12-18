package appmessage

import "github.com/Metchain/Metblock/external"

func DomainBlockToRPCBlock(block *external.DomainBlock) *RPCBlock {
	// update this now Metchainupdate
	parents := make([]*RPCBlockLevelParents, len(block.Header.DirectParents()))
	for k, v := range block.Header.ParentByteToString() {
		parents[k] = &RPCBlockLevelParents{ParentHashes: []string{v}}
	}
	header := &RPCBlockHeader{
		Version:              uint32(block.Header.Version()),
		Parents:              parents,
		HashMerkleRoot:       block.Header.Merkleroot().String(),
		AcceptedIDMerkleRoot: block.Header.Merkleroot().String(),
		UTXOCommitment:       string(block.Header.UtxoCommitment()),
		Timestamp:            block.Header.TimeInMilliseconds(),
		Bits:                 uint32(block.Header.Bits()),
		Nonce:                block.Header.Nonce(),
	}

	return &RPCBlock{
		Header: header,
	}
}
