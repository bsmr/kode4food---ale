package builtin

import (
	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/internal/async"
	"github.com/kode4food/ale/internal/stream"
)

const (
	// ChannelType is the type name for a channel
	ChannelType = data.String("channel")

	// EmitKey is the key used to emit to a Channel
	EmitKey = data.Keyword("emit")

	// SequenceKey is the key used to retrieve the Sequence from a Channel
	SequenceKey = data.Keyword("seq")
)

// Go runs the provided function asynchronously
func Go(args ...data.Value) data.Value {
	fn := args[0].(data.Caller)
	restArgs := args[1:]
	go fn.Call()(restArgs...)
	return data.Nil
}

// Chan instantiates a new go channel
func Chan(args ...data.Value) data.Value {
	var size int
	if len(args) != 0 {
		size = int(args[0].(data.Integer))
	}
	e, s := stream.NewChannel(size)

	return data.Object{
		data.TypeKey: ChannelType,
		EmitKey:      bindWriter(e),
		CloseKey:     bindCloser(e),
		SequenceKey:  s,
	}
}

// Promise instantiates a new eventually-fulfilled promise
func Promise(args ...data.Value) data.Value {
	resolver := args[0].(data.Caller).Call()
	return async.NewPromise(resolver)
}

// IsPromise returns whether the specified value is a promise
func IsPromise(args ...data.Value) data.Value {
	_, ok := args[0].(async.Promise)
	return data.Bool(ok)
}

// IsResolved returns whether the specified promise has been resolved
func IsResolved(args ...data.Value) data.Value {
	p := args[0].(async.Promise)
	return data.Bool(p.IsResolved())
}
