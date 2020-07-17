// Copyright 2020 morgine.com. All rights reserved.

package convert

import "strconv"

type Uint uint

func (u Uint) Int() (int, error) {
	return int(u), ErrOutOfIntRange.GetError(uint64(u))
}

func (u Uint) Int8() (int8, error) {
	return int8(u), ErrOutOfInt8Range.GetError(uint64(u))
}

func (u Uint) Int16() (int16, error) {
	return int16(u), ErrOutOfInt16Range.GetError(uint64(u))
}

func (u Uint) Int32() (int32, error) {
	return int32(u), ErrOutOfInt32Range.GetError(uint64(u))
}

func (u Uint) Int64() (int64, error) {
	return int64(u), ErrOutOfInt64Range.GetError(uint64(u))
}

func (u Uint) Uint() (uint, error) {
	return uint(u), nil
}

func (u Uint) Uint8() (uint8, error) {
	return uint8(u), ErrOutOfUint8Range.GetError(uint64(u))
}

func (u Uint) Uint16() (uint16, error) {
	return uint16(u), ErrOutOfUint16Range.GetError(uint64(u))
}

func (u Uint) Uint32() (uint32, error) {
	return uint32(u), ErrOutOfUint32Range.GetError(uint64(u))
}

func (u Uint) Uint64() (uint64, error) {
	return uint64(u), ErrOutOfUint64Range.GetError(uint64(u))
}

func (u Uint) Bool() (bool, error) {
	if u == 1 {
		return true, nil
	} else if u == 0 {
		return false, nil
	} else {
		return false, &BoolFormatError{
			Type:  "uint",
			Value: u,
		}
	}
}

func (u Uint) Float32() (float32, error) {
	return float32(u), nil
}

func (u Uint) Float64() (float64, error) {
	return float64(u), nil
}

func (u Uint) String() string {
	return strconv.FormatUint(uint64(u), 10)
}
