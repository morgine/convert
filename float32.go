// Copyright 2020 morgine.com. All rights reserved.

package convert

import "strconv"

type Float32 float32

func (f Float32) Int() (int, error) {
	return int(f), ErrOutOfIntRange.GetError(float32(f))
}

func (f Float32) Int8() (int8, error) {
	return int8(f), ErrOutOfInt8Range.GetError(float32(f))
}

func (f Float32) Int16() (int16, error) {
	return int16(f), ErrOutOfInt16Range.GetError(float32(f))
}

func (f Float32) Int32() (int32, error) {
	return int32(f), ErrOutOfInt32Range.GetError(float32(f))
}

func (f Float32) Int64() (int64, error) {
	return int64(f), ErrOutOfInt64Range.GetError(float32(f))
}

func (f Float32) Uint() (uint, error) {
	return uint(f), ErrOutOfUintRange.GetError(float32(f))
}

func (f Float32) Uint8() (uint8, error) {
	return uint8(f), ErrOutOfUint8Range.GetError(float32(f))
}

func (f Float32) Uint16() (uint16, error) {
	return uint16(f), ErrOutOfUint16Range.GetError(float32(f))
}

func (f Float32) Uint32() (uint32, error) {
	return uint32(f), ErrOutOfUint32Range.GetError(float32(f))
}

func (f Float32) Uint64() (uint64, error) {
	return uint64(f), ErrOutOfUint64Range.GetError(float32(f))
}

func (f Float32) Bool() (bool, error) {
	if f == 1 {
		return true, nil
	} else if f == 0 {
		return false, nil
	} else {
		return false, &BoolFormatError{
			Type:  "float32",
			Value: f,
		}
	}
}

func (f Float32) Float32() (float32, error) {
	return float32(f), nil
}

func (f Float32) Float64() (float64, error) {
	return float64(f), nil
}

func (f Float32) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
