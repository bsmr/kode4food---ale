package builtin_test

import (
	"testing"

	"github.com/kode4food/ale/data"
	"github.com/kode4food/ale/internal/assert"
	. "github.com/kode4food/ale/internal/assert/helpers"
)

func TestListEval(t *testing.T) {
	as := assert.New(t)
	as.EvalTo(`(list? '(1 2 3))`, data.True)
	as.EvalTo(`(list? '())`, data.True)
	as.EvalTo(`(list? [1 2 3])`, data.False)
	as.EvalTo(`(list? 42)`, data.False)
	as.EvalTo(`(list? (list 1 2 3))`, data.True)
	as.EvalTo(`(list)`, data.EmptyList)

	as.EvalTo(`
		(define x '(1 2 3 4))
		(x 2)
	`, F(3))
}
