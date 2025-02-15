package encoder_test

import (
	"testing"

	"github.com/kode4food/ale/internal/assert"
	"github.com/kode4food/ale/runtime/isa"
)

func TestLabels(t *testing.T) {
	as := assert.New(t)

	e := assert.GetTestEncoder()
	l1 := e.NewLabel()
	l2 := e.NewLabel()
	e.Emit(isa.Jump, l2)
	e.Emit(isa.NoOp)
	e.Emit(isa.Jump, l1)
	l2.DropAnchor()
	e.Emit(isa.NoOp)
	l1.DropAnchor()

	as.Instructions(isa.Instructions{
		isa.New(isa.Jump, 1),
		isa.New(isa.NoOp),
		isa.New(isa.Jump, 0),
		isa.New(isa.Label, 1),
		isa.New(isa.NoOp),
		isa.New(isa.Label, 0),
	}, e.Code())
}

func TestLabelDoubleAnchor(t *testing.T) {
	as := assert.New(t)

	e := assert.GetTestEncoder()
	l1 := e.NewLabel()
	l1.DropAnchor()

	defer func() {
		if rec := recover(); rec == nil {
			as.Fail("error not raised on double anchoring")
		}
	}()
	l1.DropAnchor()
}
