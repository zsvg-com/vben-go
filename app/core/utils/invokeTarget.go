package utils

import (
	"errors"
	"reflect"
)

func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	_, found := m[name]
	if !found {
		return nil, errors.New("map do not contains key ...")
	}
	fv := reflect.ValueOf(m[name])
	if fv.Kind() != reflect.Func {
		return nil, errors.New("the value of key is not a function")
	}

	if len(params) != fv.Type().NumIn() {
		return nil, errors.New("argument passed in does not match the function")
	}

	in := make([]reflect.Value, len(params))

	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}
	return fv.Call(in), nil
}
