// Copyright 2020 morgine.com. All rights reserved.

package convert

import (
	"fmt"
)

var (
	ErrOutOfIntRange    = &IntegerRangeError{Type: "int", Min: int64(MinInt), Max: uint64(MaxInt)}
	ErrOutOfInt8Range   = &IntegerRangeError{Type: "int8", Min: int64(MinInt8), Max: uint64(MaxInt8)}
	ErrOutOfInt16Range  = &IntegerRangeError{Type: "int16", Min: int64(MinInt16), Max: uint64(MaxInt16)}
	ErrOutOfInt32Range  = &IntegerRangeError{Type: "int32", Min: int64(MinInt32), Max: uint64(MaxInt32)}
	ErrOutOfInt64Range  = &IntegerRangeError{Type: "int64", Min: int64(MinInt64), Max: uint64(MaxInt64)}
	ErrOutOfUintRange   = &IntegerRangeError{Type: "uint", Min: int64(0), Max: uint64(MaxUint)}
	ErrOutOfUint8Range  = &IntegerRangeError{Type: "uint8", Min: int64(0), Max: uint64(MaxUint8)}
	ErrOutOfUint16Range = &IntegerRangeError{Type: "uint16", Min: int64(0), Max: uint64(MaxUint16)}
	ErrOutOfUint32Range = &IntegerRangeError{Type: "uint32", Min: int64(0), Max: uint64(MaxUint32)}
	ErrOutOfUint64Range = &IntegerRangeError{Type: "uint64", Min: int64(0), Max: uint64(MaxUint64)}
)

type IntegerRangeError struct {
	Type string
	Min  int64
	Max  uint64
}

func (err *IntegerRangeError) Error() string {
	return fmt.Sprintf("out of %s range %d to %d", err.Type, err.Min, err.Max)
}

func (err *IntegerRangeError) GetError(v interface{}) error {
	switch vue := v.(type) {
	case int:
		return err.getIntError(int64(vue))
	case int8:
		return err.getIntError(int64(vue))
	case int16:
		return err.getIntError(int64(vue))
	case int32:
		return err.getIntError(int64(vue))
	case int64:
		return err.getIntError(vue)
	case uint:
		return err.getUIntError(uint64(vue))
	case uint8:
		return err.getUIntError(uint64(vue))
	case uint16:
		return err.getUIntError(uint64(vue))
	case uint32:
		return err.getUIntError(uint64(vue))
	case uint64:
		return err.getUIntError(vue)
	case float32:
		return err.getFloatError(float64(vue))
	case float64:
		return err.getFloatError(vue)
	default:
		return nil
	}
}

func (err *IntegerRangeError) getIntError(i int64) error {
	if err.Min > i || uint64(i) > err.Max {
		return err
	} else {
		return nil
	}
}

func (err *IntegerRangeError) getFloatError(f float64) error {
	if float64(err.Min) > f || f > float64(err.Max) {
		return err
	} else {
		return nil
	}
}

func (err *IntegerRangeError) getUIntError(i uint64) error {
	if i > err.Max {
		return err
	} else {
		return nil
	}
}

type BoolFormatError struct {
	Type  string
	Value interface{}
}

func (b *BoolFormatError) Error() string {
	return fmt.Sprintf("%s %v format as bool", b.Type, b.Value)
}
