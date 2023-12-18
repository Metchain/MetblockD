package serialization

import "github.com/pkg/errors"

// errNoEncodingForType signifies that there's no encoding for the given type.
var errNoEncodingForType = errors.New("there's no encoding for this type")
