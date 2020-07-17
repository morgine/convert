// Copyright 2020 morgine.com. All rights reserved.

package convert

import "strconv"

type Int16 int16

func (i Int16) Int() (int, error) {
	return int(i), nil
}

func (i Int16) Int8() (int8, error) {
	return int8(i), ErrOutOfInt8Range.GetError(int64(i))
}

func (i Int16) Int16() (int16, error) {
	return int16(i), nil
}

func (i Int16) Int32() (int32, error) {
	return int32(i), nil
}

func (i Int16) Int64() (int64, error) {
	return int64(i), nil
}

func (i Int16) Uint() (uint, error) {
	return uint(i), ErrOutOfUintRange.GetError(int64(i))
}

func (i Int16) Uint8() (uint8, error) {
	return uint8(i), ErrOutOfUint8Range.GetError(int64(i))
}

func (i Int16) Uint16() (uint16, error) {
	return uint16(i), ErrOutOfUint16Range.GetError(int64(i))
}

func (i Int16) Uint32() (uint32, error) {
	return uint32(i), ErrOutOfUint32Range.GetError(int64(i))
}

func (i Int16) Uint64() (uint64, error) {
	return uint64(i), ErrOutOfUint64Range.GetError(int64(i))
}

func (i Int16) Bool() (bool, error) {
	if i == 1 {
		return true, nil
	} else if i == 0 {
		return false, nil
	} else {
		return false, &BoolFormatError{
			Type:  "int16",
			Value: i,
		}
	}
}

func (i Int16) Float32() (float32, error) {
	return float32(i), nil
}

func (i Int16) Float64() (float64, error) {
	return float64(i), nil
}

func (i Int16) String() string {
	return strconv.FormatInt(int64(i), 10)
}
