package vm

import (
	"github.com/kode4food/ale/compiler/encoder"
	"github.com/kode4food/ale/compiler/ir/optimize"
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/env"
	"github.com/kode4food/ale/runtime/isa"
	"github.com/kode4food/ale/types"
	"github.com/kode4food/ale/types/basic"
)

// Lambda encapsulates the initial environment of a virtual machine
type Lambda struct {
	Globals      env.Namespace
	Constants    data.Values
	Code         []isa.Word
	StackSize    int
	LocalCount   int
	ArityChecker data.ArityChecker
}

// LambdaFromEncoder instantiates a VM Lambda from the provided Encoder's
// intermediate representation
func LambdaFromEncoder(e encoder.Encoder) *Lambda {
	code := e.Code()
	optimized := optimize.Instructions(code)
	return &Lambda{
		Globals:    e.Globals(),
		Constants:  e.Constants(),
		StackSize:  e.StackSize(),
		LocalCount: e.LocalCount(),
		Code:       isa.Flatten(optimized),
	}
}

// Call allows a VM Lambda to be called for the purpose
// of instantiating a closure. This calling interface is used
// only by the compiler.
func (l *Lambda) Call(values ...data.Value) data.Value {
	return newClosure(l, values)
}

// CheckArity performs a compile-time arity check for the Function
func (l *Lambda) CheckArity(_ int) error {
	return nil
}

// Convention returns the Function's calling convention
func (l *Lambda) Convention() data.Convention {
	return data.NormalCall
}

// Type makes Lambda a typed value
func (l *Lambda) Type() types.Type {
	return basic.Lambda
}

// Equal compares this Lambda to another for equality
func (l *Lambda) Equal(v data.Value) bool {
	return l == v
}

func (l *Lambda) String() string {
	return data.DumpString(l)
}
