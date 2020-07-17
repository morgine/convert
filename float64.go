// Copyright 2020 morgine.com. All rights reserved.

package convert

import (
	"strconv"
)

type Float64 float64

func (f Float64) Int() (int, error) {
	return int(f), ErrOutOfIntRange.GetError(float64(f))
}

func (f Float64) Int8() (int8, error) {
	return int8(f), ErrOutOfInt8Range.GetError(float64(f))
}

func (f Float64) Int16() (int16, error) {
	return int16(f), ErrOutOfInt16Range.GetError(float64(f))
}

func (f Float64) Int32() (int32, error) {
	return int32(f), ErrOutOfInt32Range.GetError(float64(f))
}

func (f Float64) Int64() (int64, error) {
	return int64(f), ErrOutOfInt64Range.GetError(float64(f))
}

func (f Float64) Uint() (uint, error) {
	return uint(f), ErrOutOfUintRange.GetError(float64(f))
}

func (f Float64) Uint8() (uint8, error) {
	return uint8(f), ErrOutOfUint8Range.GetError(float64(f))
}

func (f Float64) Uint16() (uint16, error) {
	return uint16(f), ErrOutOfUint16Range.GetError(float64(f))
}

func (f Float64) Uint32() (uint32, error) {
	return uint32(f), ErrOutOfUint32Range.GetError(float64(f))
}

func (f Float64) Uint64() (uint64, error) {
	return uint64(f), ErrOutOfUint64Range.GetError(float64(f))
}

func (f Float64) Bool() (bool, error) {
	if f == 1 {
		return true, nil
	} else if f == 0 {
		return false, nil
	} else {
		return false, &BoolFormatError{
			Type:  "float64",
			Value: f,
		}
	}
}

func (f Float64) Float32() (float32, error) {
	return float32(f), nil
}

func (f Float64) Float64() (float64, error) {
	return float64(f), nil
}

func (f Float64) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
