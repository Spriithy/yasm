package obj

import (
	"errors"
)

var (
	// ErrOutOfBounds is used to denote an out of bounds access to a String or Tuple
	ErrOutOfBounds = errors.New("access out of bounds")
)
