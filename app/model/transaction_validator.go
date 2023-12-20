package model

import "github.com/Metchain/MetblockD/external"

// TransactionValidator exposes a set of validation classes, after which
// it's possible to determine whether a transaction is valid
type TransactionValidator interface {
	ValidateTransactionInIsolation(transaction *external.DomainTransaction, povDAAScore uint64) error
	ValidateTransactionInContextIgnoringUTXO(stagingArea *StagingArea, tx *external.DomainTransaction,
		povBlockHash *external.DomainHash, povBlockPastMedianTime int64) error
	ValidateTransactionInContextAndPopulateFee(stagingArea *StagingArea,
		tx *external.DomainTransaction, povBlockHash *external.DomainHash) error
	PopulateMass(transaction *external.DomainTransaction)
}
