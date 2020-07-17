// Copyright 2020 morgine.com. All rights reserved.

package convert

import "strconv"

type String string

func (s String) Int() (int, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	} else {
		return Int64(i).Int()
	}
}

func (s String) Int8() (int8, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	} else {
		return Int64(i).Int8()
	}
}

func (s String) Int16() (int16, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	} else {
		return Int64(i).Int16()
	}
}

func (s String) Int32() (int32, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	} else {
		return Int64(i).Int32()
	}
}

func (s String) Int64() (int64, error) {
	if s == "" {
		return 0, nil
	}
	if s == TRUE {
		return 1, nil
	}
	if s == FALSE {
		return 0, nil
	}
	return strconv.ParseInt(string(s), 10, 64)
}

func (s String) Uint() (uint, error) {
	ui, err := s.Uint64()
	if err != nil {
		return 0, err
	} else {
		return Uint64(ui).Uint()
	}
}

func (s String) Uint8() (uint8, error) {
	ui, err := s.Uint64()
	if err != nil {
		return 0, err
	} else {
		return Uint64(ui).Uint8()
	}
}

func (s String) Uint16() (uint16, error) {
	ui, err := s.Uint64()
	if err != nil {
		return 0, err
	} else {
		return Uint64(ui).Uint16()
	}
}

func (s String) Uint32() (uint32, error) {
	ui, err := s.Uint64()
	if err != nil {
		return 0, err
	} else {
		return Uint64(ui).Uint32()
	}
}

func (s String) Uint64() (uint64, error) {
	if s == "" {
		return 0, nil
	}
	if s == TRUE {
		return 1, nil
	}
	if s == FALSE {
		return 0, nil
	}
	return strconv.ParseUint(string(s), 10, 64)
}

func (s String) Bool() (bool, error) {
	if s == "" {
		return false, nil
	}
	return strconv.ParseBool(string(s))
}

func (s String) Float32() (float32, error) {
	f64, err := s.Float64()
	if err != nil {
		return 0, err
	} else {
		return Float64(f64).Float32()
	}
}

func (s String) Float64() (float64, error) {
	if s == "" {
		return 0, nil
	}
	if s == TRUE {
		return 1, nil
	}
	if s == FALSE {
		return 0, nil
	}
	return strconv.ParseFloat(string(s), 64)
}

func (s String) String() string {
	return string(s)
}
