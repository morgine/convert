// Copyright 2020 morgine.com. All rights reserved.

package convert

import (
	"fmt"
	"reflect"
)

type ErrTypeInvalid struct {
	t interface{}
}

func (e *ErrTypeInvalid) Error() string {
	return fmt.Sprintf("convert.GetValue: invalid type %v", e.t)
}

type Value interface {
	Int() (int, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)
	Uint() (uint, error)
	Uint8() (uint8, error)
	Uint16() (uint16, error)
	Uint32() (uint32, error)
	Uint64() (uint64, error)
	Bool() (bool, error)
	Float32() (float32, error)
	Float64() (float64, error)
	String() string
}

func MustGetValue(value interface{}) Value {
	t, err := GetValue(value)
	if err != nil {
		panic(err)
	}
	return t
}

func GetValue(value interface{}) (Value, error) {
	switch v := value.(type) {
	case int:
		return Int(v), nil
	case int8:
		return Int8(v), nil
	case int16:
		return Int16(v), nil
	case int32:
		return Int32(v), nil
	case int64:
		return Int64(v), nil
	case uint:
		return Uint(v), nil
	case uint8:
		return Uint(v), nil
	case uint16:
		return Uint16(v), nil
	case uint32:
		return Uint32(v), nil
	case uint64:
		return Uint64(v), nil
	case string:
		return String(v), nil
	case bool:
		return Bool(v), nil
	case float32:
		return Float32(v), nil
	case float64:
		return Float64(v), nil
	case Value:
		return v, nil
	default:
		return nil, &ErrTypeInvalid{t: v}
	}
}

func ToValue(toKind reflect.Kind, v interface{}) (interface{}, error) {
	value, err := GetValue(v)
	if err != nil {
		return nil, err
	}
	switch toKind {
	case reflect.String:
		return value.String(), nil
	case reflect.Int:
		return value.Int()
	case reflect.Int8:
		return value.Int8()
	case reflect.Int16:
		return value.Int16()
	case reflect.Int32:
		return value.Int32()
	case reflect.Int64:
		return value.Int64()
	case reflect.Uint:
		return value.Uint()
	case reflect.Uint8:
		return value.Uint8()
	case reflect.Uint16:
		return value.Uint16()
	case reflect.Uint32:
		return value.Uint32()
	case reflect.Uint64:
		return value.Uint64()
	case reflect.Float32:
		return value.Float32()
	case reflect.Float64:
		return value.Float64()
	case reflect.Bool:
		return value.Bool()
	default:
		return nil, fmt.Errorf("convert.ToSlice: unsupported type: %s", toKind)
	}
}
