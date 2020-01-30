package special

import (
	"errors"

	"github.com/kode4food/ale/compiler/encoder"
	"github.com/kode4food/ale/data"
)

// Error messages
const (
	errPatternsNotSupported = "patterns not yet supported"
)

// Pattern instantiates a matchable pattern
func Pattern(e encoder.Type, args ...data.Value) {
	panic(errors.New(errPatternsNotSupported))
}
