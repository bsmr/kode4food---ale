package data_test

import (
	"testing"

	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/internal/assert"
	. "github.com/kode4food/ale/internal/assert/helpers"
	"github.com/kode4food/ale/internal/sequence"
)

func TestLastOfSequence(t *testing.T) {
	as := assert.New(t)

	v, ok := data.Last(data.EmptyList)
	as.Nil(v)
	as.False(ok)

	v, ok = data.Last(L(S("this"), S("is"), S("last")))
	as.String("last", v)
	as.True(ok)

	v, ok = data.Last(V(S("this"), S("is"), S("last")))
	as.String("last", v)
	as.True(ok)

	v, ok = data.Last(sequence.NewLazy(
		func() (data.Value, data.Sequence, bool) {
			return S("hello"), data.EmptyList, true
		},
	))
	as.String("hello", v)
	as.True(ok)

	v, ok = data.Last(sequence.NewLazy(
		func() (data.Value, data.Sequence, bool) {
			return data.Nil, data.EmptyList, false
		},
	))
	as.False(ok)
}
