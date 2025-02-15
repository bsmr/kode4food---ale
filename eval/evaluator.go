package eval

import (
	"github.com/kode4food/ale/compiler/encoder"
	"github.com/kode4food/ale/compiler/generate"
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/env"
	"github.com/kode4food/ale/read"
	"github.com/kode4food/ale/runtime/isa"
	"github.com/kode4food/ale/runtime/vm"
)

// String evaluates the specified raw source
func String(ns env.Namespace, src data.String) data.Value {
	r := read.FromString(src)
	return Block(ns, r)
}

// Block evaluates a Sequence that a call to FromScanner might produce
func Block(ns env.Namespace, s data.Sequence) data.Value {
	var res data.Value
	for f, r, ok := s.Split(); ok; f, r, ok = r.Split() {
		res = Value(ns, f)
	}
	return res
}

// Value evaluates the provided Value
func Value(ns env.Namespace, v data.Value) data.Value {
	e := encoder.NewEncoder(ns)
	generate.Value(e, v)
	e.Emit(isa.Return)
	return encodeAndRun(e)
}

func encodeAndRun(e encoder.Encoder) data.Value {
	fn := vm.LambdaFromEncoder(e)
	closure := fn.Call().(data.Function)
	return closure.Call()
}
