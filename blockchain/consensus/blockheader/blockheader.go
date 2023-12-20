package blockheader

import (
	"github.com/Metchain/MetblockD/external"
)

func NewImmutableBlockHeader(
	version uint16,
	blockheight int64,
	blockhash *external.DomainHash,
	previoushash *external.DomainHash,
	merkleroot *external.DomainHash,
	parents []external.BlockLevelParents,
	metblock *external.DomainHash,
	megablock *external.DomainHash,
	childblocks []external.BlockLevelChildern,
	timeInMilliseconds int64,
	bits uint64,
	nonce uint64,
	btype int,
	verificationPoint *external.DomainHash,
	utxo []byte,
) external.BlockHeader {

	return &blockHeader{
		version:            version,
		blockheight:        blockheight,
		blockhash:          blockhash,
		previoushash:       previoushash,
		merkleroot:         merkleroot,
		parents:            parents,
		metblock:           metblock,
		megablock:          megablock,
		childblocks:        childblocks,
		timeInMilliseconds: timeInMilliseconds,

		bits:              bits,
		nonce:             nonce,
		blockLevel:        btype,
		verificationPoint: verificationPoint,
		utxocommitment:    utxo,
	}
}

type blockHeader struct {
	version            uint16
	blockhash          *external.DomainHash
	blockheight        int64
	previoushash       *external.DomainHash
	merkleroot         *external.DomainHash
	metblock           *external.DomainHash
	megablock          *external.DomainHash
	childblocks        []external.BlockLevelChildern
	parents            []external.BlockLevelParents
	timeInMilliseconds int64
	bits               uint64
	nonce              uint64
	blockLevel         int
	verificationPoint  *external.DomainHash
	bytpe              int
	utxocommitment     []byte
}

func (bh *blockHeader) ToImmutable() external.BlockHeader {
	return bh.clone()
}

func (bh *blockHeader) SetNonce(nonce uint64) {

	bh.nonce = nonce
}

func (bh *blockHeader) SetTimeInMilliseconds(timeInMilliseconds int64) {

	bh.timeInMilliseconds = timeInMilliseconds
}

func (bh *blockHeader) SetHashMerkleRoot(hashMerkleRoot *external.DomainHash) {

	bh.merkleroot = hashMerkleRoot
}
func (bh *blockHeader) UtxoCommitment() []byte {
	//TODO implement me
	return bh.utxocommitment
}

func (bh *blockHeader) Btype() int {
	//TODO implement me
	return bh.bytpe
}

func (bh *blockHeader) DirectParents() external.BlockLevelParents {
	if len(bh.parents) == 0 {
		return external.BlockLevelParents{}
	}

	return bh.parents[0]
}

func (bh *blockHeader) ChildBlocks() []external.BlockLevelChildern {
	//TODO implement me
	return bh.childblocks
}

func (bh *blockHeader) BlockLevel(maxBlockLevel int) int {
	//TODO implement me
	return bh.blockLevel
}

func (bh *blockHeader) ToMutable() external.MutableBlockHeader {
	//TODO implement me
	return bh.clone()
}

func (bh *blockHeader) Version() uint16 {
	if bh.verificationPoint.String() == "" {
		return 0
	}
	return bh.version
}

func (bh *blockHeader) Blockheight() uint64 {
	//TODO implement me
	return uint64(bh.blockheight)
}

func (bh *blockHeader) BlockHash() *external.DomainHash {
	//TODO implement me
	return bh.blockhash
}

func (bh *blockHeader) Previoushash() *external.DomainHash {
	//TODO implement me
	return bh.previoushash
}

func (bh *blockHeader) Merkleroothash() *external.DomainHash {
	//TODO implement me
	return bh.merkleroot
}

func (bh *blockHeader) MetBlock() *external.DomainHash {
	//TODO implement me
	return bh.metblock
}

func (bh *blockHeader) MegaBlock() *external.DomainHash {
	//TODO implement me
	return bh.megablock
}

func (bh *blockHeader) UTXOCommitment() *external.DomainHash {
	//TODO implement me
	panic("implement me")
}

func (bh *blockHeader) Bits() uint64 {
	//TODO implement me
	return bh.bits
}

func (bh *blockHeader) TimeInMilliseconds() int64 {
	return bh.timeInMilliseconds
}

func (bh *blockHeader) Nonce() uint64 {
	return bh.nonce
}
func (bc *blockHeader) Parents() []external.BlockLevelParents {
	return bc.parents
}

func (bh *blockHeader) Equal(other external.BaseBlockHeader) bool {
	/*if bh == nil || other == nil {
		return bh == other
	}

	// If only the underlying value of other is nil it'll
	// make `other == nil` return false, so we check it
	// explicitly.
	downcastedOther := other.(*blockHeader)
	if bh == nil || downcastedOther == nil {
		return bh == downcastedOther
	}

	if bh.version != other.Version() {
		return false
	}

	if !bh.hashMerkleRoot.Equal(other.HashMerkleRoot()) {
		return false
	}

	if !bh.acceptedIDMerkleRoot.Equal(other.AcceptedIDMerkleRoot()) {
		return false
	}

	if !bh.utxoCommitment.Equal(other.UTXOCommitment()) {
		return false
	}

	if bh.timeInMilliseconds != other.TimeInMilliseconds() {
		return false
	}

	if bh.bits != other.Bits() {
		return false
	}

	if bh.nonce != other.Nonce() {
		return false
	}*/

	return true
}

func (bh *blockHeader) clone() *blockHeader {
	return &blockHeader{
		version:            bh.version,
		blockheight:        bh.blockheight,
		blockhash:          bh.blockhash,
		previoushash:       bh.previoushash,
		merkleroot:         bh.merkleroot,
		parents:            bh.parents,
		metblock:           bh.metblock,
		megablock:          bh.megablock,
		childblocks:        bh.childblocks,
		timeInMilliseconds: bh.timeInMilliseconds,

		bits:              bh.bits,
		nonce:             bh.nonce,
		blockLevel:        bh.blockLevel,
		verificationPoint: bh.verificationPoint,
		utxocommitment:    bh.utxocommitment,
	}
}
