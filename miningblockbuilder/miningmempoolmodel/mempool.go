package miningmempoolmodel

import "github.com/Metchain/MetblockD/external"

type Mempool interface {
	HandleNewBlockTransactions(txs []*external.DomainTransaction) ([]*external.DomainTransaction, error)
	BlockCandidateTransactions() []*external.DomainTransaction
	ValidateAndInsertTransaction(transaction *external.DomainTransaction, isHighPriority bool, allowOrphan bool) (
		acceptedTransactions []*external.DomainTransaction, err error)
	RemoveTransactions(txs []*external.DomainTransaction, removeRedeemers bool) error
	GetTransaction(
		transactionID *external.DomainTransactionID,
		includeTransactionPool bool,
		includeOrphanPool bool,
	) (
		transactionPoolTransaction *external.DomainTransaction,
		isOrphan bool,
		found bool)
	GetTransactionsByAddresses(
		includeTransactionPool bool,
		includeOrphanPool bool) (
		sendingInTransactionPool map[string]*external.DomainTransaction,
		receivingInTransactionPool map[string]*external.DomainTransaction,
		sendingInOrphanPool map[string]*external.DomainTransaction,
		receivingInOrphanPool map[string]*external.DomainTransaction,
		err error)
	AllTransactions(
		includeTransactionPool bool,
		includeOrphanPool bool,
	) (
		transactionPoolTransactions []*external.DomainTransaction,
		orphanPoolTransactions []*external.DomainTransaction)
	TransactionCount(
		includeTransactionPool bool,
		includeOrphanPool bool) int

	//Keep this for layer 3
	//RevalidateHighPriorityTransactions() (validTransactions []*external.DomainTransaction, err error)
}
