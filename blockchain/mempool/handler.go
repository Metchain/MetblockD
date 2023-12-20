package mempool

import "github.com/Metchain/MetblockD/external"

func (m mempool) HandleNewBlockTransactions(txs []*external.DomainTransaction) ([]*external.DomainTransaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m mempool) BlockCandidateTransactions() []*external.DomainTransaction {
	//TODO implement me
	panic("implement me")
}

func (m mempool) ValidateAndInsertTransaction(transaction *external.DomainTransaction, isHighPriority bool, allowOrphan bool) (acceptedTransactions []*external.DomainTransaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mempool) RemoveTransactions(txs []*external.DomainTransaction, removeRedeemers bool) error {
	//TODO implement me
	panic("implement me")
}

func (m mempool) GetTransaction(transactionID *external.DomainTransactionID, includeTransactionPool bool, includeOrphanPool bool) (transactionPoolTransaction *external.DomainTransaction, isOrphan bool, found bool) {
	//TODO implement me
	panic("implement me")
}

func (m mempool) GetTransactionsByAddresses(includeTransactionPool bool, includeOrphanPool bool) (sendingInTransactionPool map[string]*external.DomainTransaction, receivingInTransactionPool map[string]*external.DomainTransaction, sendingInOrphanPool map[string]*external.DomainTransaction, receivingInOrphanPool map[string]*external.DomainTransaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mempool) AllTransactions(includeTransactionPool bool, includeOrphanPool bool) (transactionPoolTransactions []*external.DomainTransaction, orphanPoolTransactions []*external.DomainTransaction) {
	//TODO implement me
	panic("implement me")
}

func (m mempool) TransactionCount(includeTransactionPool bool, includeOrphanPool bool) int {
	//TODO implement me
	panic("implement me")
}
