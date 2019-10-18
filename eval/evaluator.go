package eval

import (
	"github.com/kode4food/ale/compiler/encoder"
	"github.com/kode4food/ale/compiler/generate"
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/namespace"
	"github.com/kode4food/ale/read"
	"github.com/kode4food/ale/runtime/isa"
	"github.com/kode4food/ale/runtime/vm"
)

// String evaluates the specified raw source
func String(ns namespace.Type, src data.String) data.Value {
	r := read.FromString(src)
	return Block(ns, r)
}

// Block evaluates a Sequence that a call to FromScanner might produce
func Block(ns namespace.Type, s data.Sequence) data.Value {
	var res data.Value
	for f, r, ok := s.Split(); ok; f, r, ok = r.Split() {
		e := encoder.NewEncoder(ns)
		generate.Value(e, f)
		e.Emit(isa.Return)
		res = encodeAndRun(e)
	}
	return res
}

// Value evaluates the provided Value
func Value(ns namespace.Type, v data.Value) data.Value {
	e := encoder.NewEncoder(ns)
	generate.Value(e, v)
	e.Emit(isa.Return)
	return encodeAndRun(e)
}

func encodeAndRun(e encoder.Type) data.Value {
	fn := vm.LambdaFromEncoder(e).Call()
	closure := fn().(data.Caller).Call()
	return closure()
}
