package ffi

import (
	"reflect"

	"github.com/kode4food/ale/data"
)

type (
	stringWrapper reflect.Kind
	boolWrapper   bool
)

var (
	_stringWrapper stringWrapper
	_boolWrapper   boolWrapper

	stringZero = reflect.ValueOf("")
	boolZero   = reflect.ValueOf(false)
)

func makeWrappedBool(_ reflect.Type) (Wrapper, error) {
	return _boolWrapper, nil
}

func makeWrappedString(_ reflect.Type) (Wrapper, error) {
	return _stringWrapper, nil
}

func (stringWrapper) Wrap(_ *Context, v reflect.Value) (data.Value, error) {
	return data.String(v.Interface().(string)), nil
}

func (stringWrapper) Unwrap(v data.Value) (reflect.Value, error) {
	if v == nil {
		return stringZero, nil
	}
	return reflect.ValueOf(v.String()), nil
}

func (b boolWrapper) Wrap(_ *Context, v reflect.Value) (data.Value, error) {
	return data.Bool(v.Bool()), nil
}

func (b boolWrapper) Unwrap(v data.Value) (reflect.Value, error) {
	if v == nil {
		return boolZero, nil
	}
	return reflect.ValueOf(bool(v.(data.Bool))), nil
}