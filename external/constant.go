package external

const (
	// StatusInvalid indicates that the block is invalid.
	StatusInvalid BlockStatus = iota

	// StatusUTXOValid indicates the block is valid from any UTXO related aspects and has passed all the other validations as well.
	StatusUTXOValid

	// StatusUTXOPendingVerification indicates that the block is pending verification against its past UTXO-Set, either
	// because it was not yet verified since the block was never in the selected parent chain, or if the
	// block violates finality.
	StatusUTXOPendingVerification

	// StatusDisqualifiedFromChain indicates that the block is not eligible to be a selected parent.
	StatusDisqualifiedFromChain

	// StatusHeaderOnly indicates that the block transactions are not held (pruned or wasn't added yet)
	StatusHeaderOnly
)
