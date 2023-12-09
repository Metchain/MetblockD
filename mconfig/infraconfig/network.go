package infraconfig

import (
	"fmt"
	"github.com/Metchain/Metblock/mconfig/dagconfig"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"os"
)

// ResolveNetwork parses the network command line argument and sets NetParams accordingly.
// It returns error if more than one network was selected, nil otherwise.
func (networkFlags *NetworkFlags) ResolveNetwork(parser *flags.Parser) error {
	//NetParams holds the selected network parameters. Default value is main-net.
	networkFlags.ActiveNetParams = &dagconfig.MainnetParams
	// Multiple networks can't be selected simultaneously.
	numNets := 0

	if numNets > 1 {
		message := "Multiple networks parameters (testnet, simnet, devnet, etc.) cannot be used" +
			"together. Please choose only one network"
		err := errors.Errorf(message)
		fmt.Fprintln(os.Stderr, err)
		parser.WriteHelp(os.Stderr)
		return err
	}

	return nil
}

// NetParams returns the ActiveNetParams
func (networkFlags *NetworkFlags) NetParams() *dagconfig.Params {
	return networkFlags.ActiveNetParams
}
