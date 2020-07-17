// Copyright 2020 morgine.com. All rights reserved.

package convert

import "strconv"

type Uint8 uint8

func (u Uint8) Int() (int, error) {
	return int(u), nil
}

func (u Uint8) Int8() (int8, error) {
	return int8(u), ErrOutOfInt8Range.GetError(uint64(u))
}

func (u Uint8) Int16() (int16, error) {
	return int16(u), nil
}

func (u Uint8) Int32() (int32, error) {
	return int32(u), nil
}

func (u Uint8) Int64() (int64, error) {
	return int64(u), nil
}

func (u Uint8) Uint() (uint, error) {
	return uint(u), nil
}

func (u Uint8) Uint8() (uint8, error) {
	return uint8(u), nil
}

func (u Uint8) Uint16() (uint16, error) {
	return uint16(u), nil
}

func (u Uint8) Uint32() (uint32, error) {
	return uint32(u), nil
}

func (u Uint8) Uint64() (uint64, error) {
	return uint64(u), nil
}

func (u Uint8) Bool() (bool, error) {
	if u == 1 {
		return true, nil
	} else if u == 0 {
		return false, nil
	} else {
		return false, &BoolFormatError{
			Type:  "uint8",
			Value: u,
		}
	}
}

func (u Uint8) Float32() (float32, error) {
	return float32(u), nil
}

func (u Uint8) Float64() (float64, error) {
	return float64(u), nil
}

func (u Uint8) String() string {
	return strconv.FormatUint(uint64(u), 10)
}
