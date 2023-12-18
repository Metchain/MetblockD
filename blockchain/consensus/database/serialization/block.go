package serialization

import "github.com/Metchain/Metblock/external"

// DbBlockToDomainBlock converts DbBlock to DomainBlock
func DbBlockToDomainBlock(dbBlock *DbBlock) (*external.DomainBlock, error) {
	domainBlockHeader, err := DbBlockHeaderToDomainBlockHeader(dbBlock.Header)
	if err != nil {
		return nil, err
	}

	domainTransactions := make([]*external.DomainTransaction, len(dbBlock.Transactions))

	return &external.DomainBlock{
		Header:       domainBlockHeader,
		Transactions: domainTransactions,
	}, nil
}
