package controller

import "github.com/Metchain/Metblock/external"

// IDToTransactionMap maps transactionID to a MempoolTransaction
type IDToTransactionMap map[external.DomainTransactionID]*MempoolTransaction

// IDToTransactionsSliceMap maps transactionID to a slice MempoolTransaction
type IDToTransactionsSliceMap map[external.DomainTransactionID][]*MempoolTransaction

// OutpointToUTXOEntryMap maps an outpoint to a UTXOEntry
type OutpointToUTXOEntryMap map[external.DomainOutpoint]external.UTXOEntry

// OutpointToTransactionMap maps an outpoint to a MempoolTransaction
type OutpointToTransactionMap map[external.DomainOutpoint]*MempoolTransaction

// ScriptPublicKeyStringToDomainTransaction maps an outpoint to a DomainTransaction
type ScriptPublicKeyStringToDomainTransaction map[string]*external.DomainTransaction

// OrphanTransaction represents a transaction in the OrphanPool
type RejectedTransaction struct {
	transaction *external.DomainTransaction
}
