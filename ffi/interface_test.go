package ffi_test

import (
	"testing"

	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/ffi"
	"github.com/kode4food/ale/internal/assert"
	. "github.com/kode4food/ale/internal/assert/helpers"
)

type (
	testInterface interface {
		Void(func())
		Add(int, int) int
		Double(int, int) (int, int)
	}

	testReceiver bool
)

func TestVoidInterface(t *testing.T) {
	as := assert.New(t)
	f := makeCall(ffi.Wrap(func() testInterface {
		return testReceiver(false)
	}))
	r := f().(data.Object)
	as.Equal(3, len(r))

	b := []bool{false}
	m := r[K("Void")].(data.Call)
	m(ffi.Wrap(func() {
		b[0] = true
	}))
	as.True(b[0])
}

func TestValueInterface(t *testing.T) {
	as := assert.New(t)
	f := makeCall(ffi.Wrap(func() testInterface {
		return testReceiver(false)
	}))
	r := f().(data.Object)
	as.Equal(3, len(r))

	m := r[K("Add")].(data.Call)
	s := m(ffi.Wrap(I(4)), ffi.Wrap(I(6)))
	as.Equal(I(10), s)
}

func TestVectorInterface(t *testing.T) {
	as := assert.New(t)
	f := makeCall(ffi.Wrap(func() testInterface {
		return testReceiver(false)
	}))
	r := f().(data.Object)
	as.Equal(3, len(r))

	m := r[K("Double")].(data.Call)
	d := m(ffi.Wrap(I(4)), ffi.Wrap(I(6))).(data.Vector)
	as.Equal(2, len(d))
	as.Equal(I(8), d[0])
	as.Equal(I(12), d[1])
}

func (testReceiver) Void(f func()) {
	f()
}

func (testReceiver) Add(l, r int) int {
	return l + r
}

func (testReceiver) Double(f, s int) (int, int) {
	return f * 2, s * 2
}
