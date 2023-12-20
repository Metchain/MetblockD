package blockbuilder

import (
	"errors"
	"github.com/Metchain/MetblockD/app/model"
	"github.com/Metchain/MetblockD/blockchain/consensus/ruleerrors"
	"github.com/Metchain/MetblockD/external"
)

func (bb *blockBuilder) validateTransactions(stagingArea *model.StagingArea,
	transactions []*external.DomainTransaction) error {

	invalidTransactions := make([]ruleerrors.InvalidTransaction, 0)
	for _, transaction := range transactions {
		err := bb.validateTransaction(stagingArea, transaction)
		if err != nil {
			if !errors.As(err, &ruleerrors.RuleError{}) {
				return err
			}
			invalidTransactions = append(invalidTransactions,
				ruleerrors.InvalidTransaction{Transaction: transaction, Error: err})
		}
	}

	if len(invalidTransactions) > 0 {
		return ruleerrors.NewErrInvalidTransactionsInNewBlock(invalidTransactions)
	}

	return nil
}

func (bb *blockBuilder) validateTransaction(
	stagingArea *model.StagingArea, transaction *external.DomainTransaction) error {

	originalEntries := make([]external.UTXOEntry, len(transaction.Inputs))
	for i, input := range transaction.Inputs {
		originalEntries[i] = input.UTXOEntry
		input.UTXOEntry = nil
	}

	defer func() {
		for i, input := range transaction.Inputs {
			input.UTXOEntry = originalEntries[i]
		}
	}()
	return nil
	/*err := bb.consensusStateManager.PopulateTransactionWithUTXOEntries(stagingArea, transaction)
	if err != nil {
		return err
	}

	virtualPastMedianTime, err := bb.pastMedianTimeManager.PastMedianTime(stagingArea, model.VirtualBlockHash)
	if err != nil {
		return err
	}

	err = bb.transactionValidator.ValidateTransactionInContextIgnoringUTXO(stagingArea, transaction, model.VirtualBlockHash, virtualPastMedianTime)
	if err != nil {
		return err
	}

	return bb.transactionValidator.ValidateTransactionInContextAndPopulateFee(stagingArea, transaction, model.VirtualBlockHash)*/
}
