package builtin

import (
	"gitlab.com/kode4food/ale/data"
	"gitlab.com/kode4food/ale/stdlib"
)

const (
	// ChannelKey is the key used to identify a Channel
	ChannelKey = data.Keyword("channel")

	// EmitKey is the key used to emit to a Channel
	EmitKey = data.Keyword("emit")

	// SequenceKey is the key used to retrieve the Sequence from a Channel
	SequenceKey = data.Keyword("seq")
)

var channelPrototype = data.Object{
	ChannelKey: data.True,
}

// Go runs the provided function asynchronously
func Go(args ...data.Value) data.Value {
	fn := args[0].(data.Caller)
	restArgs := args[1:]
	go fn.Caller()(restArgs...)
	return data.Nil
}

// Chan instantiates a new go channel
func Chan(_ ...data.Value) data.Value {
	e, s := stdlib.NewChannel()

	return channelPrototype.Extend(data.Object{
		EmitKey:     bindWriter(e),
		CloseKey:    bindCloser(e),
		SequenceKey: s,
	})
}

// Promise instantiates a new eventually-fulfilled promise
func Promise(args ...data.Value) data.Value {
	if len(args) == 0 {
		return stdlib.NewPromise()
	}
	p := stdlib.NewPromise()
	p.Deliver(args[0])
	return p
}

// IsPromise returns whether or not the specified value is a promise
func IsPromise(args ...data.Value) data.Value {
	_, ok := args[0].(stdlib.Promise)
	return data.Bool(ok)
}