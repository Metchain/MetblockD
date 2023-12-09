package util

const (
	// Unknown/Erroneous prefix
	Bech32PrefixUnknown Bech32Prefix = iota

	// Prefix for the main network.
	Bech32PrefixMet

	// Prefix for the test network.
	Bech32PrefixMetTest
)

// Map from strings to Bech32 address prefix constants for parsing purposes.
var stringsToBech32Prefixes = map[string]Bech32Prefix{
	"met":     Bech32PrefixMet,
	"mettest": Bech32PrefixMetTest,
}
