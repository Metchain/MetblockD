package external

type BlockWithTrustedData struct {
	Block       *DomainBlock
	MetblockDAA []*TrustedDataDataMETBlockHeader
	MetDagData  []*BlockMetDataHashPair
}

// TrustedDataDataDAAHeader is a block that belongs to BlockWithTrustedData.DAAWindow
type TrustedDataDataMETBlockHeader struct {
	Header      BlockHeader
	MetGDagData *MetGDagData
}

// BlockGHOSTDAGDataHashPair is a pair of a block hash and its ghostdag data
type BlockMetDataHashPair struct {
	Hash        *DomainHash
	MetGDagData *MetGDagData
}

type MetGDagData struct {
	selectedParent *DomainHash
}
