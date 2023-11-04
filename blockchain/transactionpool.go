package blockchain

func (bc *Blockchain) TransactionPool() []*Transaction {
	return bc.Txpool
}

func (bc *Blockchain) ClearTransactionPool() {
	bc.Txpool = bc.Txpool[:0]
	bc.NFTPool = bc.NFTPool[:0]
	
}
