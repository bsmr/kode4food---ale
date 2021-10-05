package basic_test

import (
	"testing"

	"github.com/kode4food/ale/types/basic"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	as := assert.New(t)

	as.Equal("any", basic.Any.Name())
	as.True(basic.Any.Accepts(basic.Lambda))
	as.True(basic.Any.Accepts(basic.Number))
	as.True(basic.Any.Accepts(basic.Any))
}