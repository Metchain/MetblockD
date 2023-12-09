package execenv

import (
	"fmt"
	"github.com/Metchain/Metblock/utils/limits"
	"os"
	"runtime"
)

// Initialize initializes the execution environment required to run Metchaind
func Initialize(desiredLimits *limits.DesiredLimits) {
	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Up some limits.
	if err := limits.SetLimits(desiredLimits); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set limits: %s\n", err)
		os.Exit(1)
	}

}
