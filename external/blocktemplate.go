package external

// DomainBlockTemplate contains a Block plus metadata related to its generation
type DomainBlockTemplate struct {
	Block        *DomainBlock
	CoinbaseData *DomainCoinbaseData

	IsNearlySynced bool
}
