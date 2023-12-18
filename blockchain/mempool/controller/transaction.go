package controller

import "github.com/Metchain/Metblock/external"

// MempoolTransaction represents a transaction inside the main TransactionPool
type MempoolTransaction struct {
	transaction              *external.DomainTransaction
	parentTransactionsInPool IDToTransactionMap
}

// NewMempoolTransaction constructs a new MempoolTransaction
func NewMempoolTransaction(
	transaction *external.DomainTransaction,
	parentTransactionsInPool IDToTransactionMap) *MempoolTransaction {
	return &MempoolTransaction{
		transaction:              transaction,
		parentTransactionsInPool: parentTransactionsInPool,
	}
}
