package addressmanager

import (
	"errors"
	"github.com/Metchain/MetblockD/utils/logger"
)

// ErrAddressNotFound is an error returned from some functions when a
// given address is not found in the address manager
var ErrAddressNotFound = errors.New("address not found")

var log = logger.RegisterSubSystem("METD-Address-Manager")
