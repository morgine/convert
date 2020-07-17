// Copyright 2020 morgine.com. All rights reserved.

package convert

const (
	TRUE  = "true"
	FALSE = "false"
)

type Bool bool

func (b Bool) Int() (int, error) {
	return int(b.getInt8()), nil
}

func (b Bool) getInt8() int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func (b Bool) Int8() (int8, error) {
	return b.getInt8(), nil
}

func (b Bool) Int16() (int16, error) {
	return int16(b.getInt8()), nil
}

func (b Bool) Int32() (int32, error) {
	return int32(b.getInt8()), nil
}

func (b Bool) Int64() (int64, error) {
	return int64(b.getInt8()), nil
}

func (b Bool) Uint() (uint, error) {
	return uint(b.getInt8()), nil
}

func (b Bool) Uint8() (uint8, error) {
	return uint8(b.getInt8()), nil
}

func (b Bool) Uint16() (uint16, error) {
	return uint16(b.getInt8()), nil
}

func (b Bool) Uint32() (uint32, error) {
	return uint32(b.getInt8()), nil
}

func (b Bool) Uint64() (uint64, error) {
	return uint64(b.getInt8()), nil
}

func (b Bool) Bool() (bool, error) {
	return bool(b), nil
}

func (b Bool) Float32() (float32, error) {
	return float32(b.getInt8()), nil
}

func (b Bool) Float64() (float64, error) {
	return float64(b.getInt8()), nil
}

func (b Bool) String() string {
	if b {
		return TRUE
	} else {
		return FALSE
	}
}
