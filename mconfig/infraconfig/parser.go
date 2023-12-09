package infraconfig

import (
	"github.com/jessevdk/go-flags"
	"runtime"
)

// newConfigParser returns a new command line flags parser.
func newConfigParser(cfgFlags *Flags, options flags.Options) *flags.Parser {
	parser := flags.NewParser(cfgFlags, options)
	if runtime.GOOS == "windows" {
		parser.AddGroup("Service Options", "Service Options", cfgFlags.ServiceOptions)
	}
	return parser
}
