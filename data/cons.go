package data

import (
	"bytes"
)

type (
	// Pair represents the interface for a binary structure, such as a Cons
	Pair interface {
		Value
		Car() Value
		Cdr() Value
	}

	// Pairs represents multiple pairs
	Pairs []Pair

	// Cons cells are the standard implementation of a Pair
	Cons [2]Value
)

// Cons fields
const (
	Car = 0
	Cdr = 1
)

// NewCons returns a new Cons cell instance
func NewCons(car, cdr Value) Cons {
	return Cons{car, cdr}
}

// Car returns the first element of a Pair
func (c Cons) Car() Value {
	return c[Car]
}

// Cdr returns the second element of a Pair
func (c Cons) Cdr() Value {
	return c[Cdr]
}

// Equal compares this Cons to another for equality
func (c Cons) Equal(v Value) bool {
	if v, ok := v.(Cons); ok {
		return c == v || c[Car].Equal(v[Car]) && c[Cdr].Equal(v[Cdr])
	}
	return false
}

func (c Cons) String() string {
	var buf bytes.Buffer
	buf.WriteByte('(')
	var next Pair = c
	for {
		buf.WriteString(MaybeQuoteString(next.Car()))
		cdr := next.Cdr()
		if s, ok := cdr.(Sequence); ok && s.IsEmpty() {
			break
		} else if p, ok := cdr.(Pair); ok {
			buf.WriteByte(' ')
			next = p
		} else {
			buf.WriteString(" . ")
			buf.WriteString(MaybeQuoteString(cdr))
			break
		}
	}
	buf.WriteByte(')')
	return buf.String()
}

// HashCode returns the has code for this Cons
func (c Cons) HashCode() uint64 {
	return HashCode(c[Car]) ^ HashCode(c[Cdr])
}
