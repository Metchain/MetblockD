package util

// Bech32Prefix is the human-readable prefix for a Bech32 address.
type Bech32Prefix int

type Address interface {
	// String returns the string encoding of the transaction output
	// destination.
	//
	// Please note that String differs subtly from EncodeAddress: String
	// will return the value as a string without any conversion, while
	// EncodeAddress may convert destination types (for example,
	// converting pubkeys to P2PK addresses) before encoding as a
	// payment address string.
	String() string

	// EncodeAddress returns the string encoding of the payment address
	// associated with the Address value. See the comment on String
	// for how this method differs from String.
	EncodeAddress() string

	// ScriptAddress returns the raw bytes of the address to be used
	// when inserting the address into a txout's script.
	ScriptAddress() []byte

	// Prefix returns the prefix for this address
	Prefix() Bech32Prefix

	// IsForPrefix returns whether or not the address is associated with the
	// passed metchain network.
	IsForPrefix(prefix Bech32Prefix) bool
}
