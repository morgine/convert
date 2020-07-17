// Copyright 2020 morgine.com. All rights reserved.

package convert

import (
	"fmt"
	"reflect"
)

type Slice struct {
	vs   []Value
	i    []int
	i8   []int8
	i16  []int16
	i32  []int32
	i64  []int64
	ui   []uint
	ui8  []uint8
	ui16 []uint16
	ui32 []uint32
	ui64 []uint64
	str  []string
	f64  []float64
	f32  []float32
	b    []bool
}

func ReflectToSlice(v interface{}) []interface{} {
	rv := reflect.Indirect(reflect.ValueOf(v))
	if rv.Kind() == reflect.Slice {
		length := rv.Len()
		vs := make([]interface{}, length)
		for i := 0; i < length; i++ {
			vs[i] = rv.Index(i).Interface()
		}
		return vs
	}
	return nil
}

func NewSlice(v []interface{}) (*Slice, error) {
	vs := make([]Value, len(v))
	for i, elem := range v {
		value, err := GetValue(elem)
		if err != nil {
			return nil, err
		}
		vs[i] = value
	}
	return &Slice{vs: vs}, nil
}

func (s *Slice) Int() ([]int, error) {
	if s.i == nil {
		vs := make([]int, len(s.vs))
		for idx, elem := range s.vs {
			i, err := elem.Int()
			if err != nil {
				return nil, err
			}
			vs[idx] = i
		}
		s.i = vs
	}
	return s.i, nil
}

func (s *Slice) Int8() ([]int8, error) {
	if s.i8 == nil {
		vs := make([]int8, len(s.vs))
		for idx, elem := range s.vs {
			i8, err := elem.Int8()
			if err != nil {
				return nil, err
			}
			vs[idx] = i8
		}
		s.i8 = vs
	}
	return s.i8, nil
}

func (s *Slice) Int16() ([]int16, error) {
	if s.i16 == nil {
		vs := make([]int16, len(s.vs))
		for idx, elem := range s.vs {
			i16, err := elem.Int16()
			if err != nil {
				return nil, err
			}
			vs[idx] = i16
		}
		s.i16 = vs
	}
	return s.i16, nil
}

func (s *Slice) Int32() ([]int32, error) {
	if s.i32 == nil {
		vs := make([]int32, len(s.vs))
		for idx, elem := range s.vs {
			i, err := elem.Int32()
			if err != nil {
				return nil, err
			}
			vs[idx] = i
		}
		s.i32 = vs
	}
	return s.i32, nil
}

func (s *Slice) Int64() ([]int64, error) {
	if s.i64 == nil {
		vs := make([]int64, len(s.vs))
		for idx, elem := range s.vs {
			i64, err := elem.Int64()
			if err != nil {
				return nil, err
			}
			vs[idx] = i64
		}
		s.i64 = vs
	}
	return s.i64, nil
}

func (s *Slice) Uint() ([]uint, error) {
	if s.ui == nil {
		vs := make([]uint, len(s.vs))
		for idx, elem := range s.vs {
			ui, err := elem.Uint()
			if err != nil {
				return nil, err
			}
			vs[idx] = ui
		}
		s.ui = vs
	}
	return s.ui, nil
}

func (s *Slice) Uint8() ([]uint8, error) {
	if s.ui8 == nil {
		vs := make([]uint8, len(s.vs))
		for idx, elem := range s.vs {
			ui8, err := elem.Uint8()
			if err != nil {
				return nil, err
			}
			vs[idx] = ui8
		}
		s.ui8 = vs
	}
	return s.ui8, nil
}

func (s *Slice) Uint16() ([]uint16, error) {
	if s.ui16 == nil {
		vs := make([]uint16, len(s.vs))
		for idx, elem := range s.vs {
			ui16, err := elem.Uint16()
			if err != nil {
				return nil, err
			}
			vs[idx] = ui16
		}
		s.ui16 = vs
	}
	return s.ui16, nil
}

func (s *Slice) Uint32() ([]uint32, error) {
	if s.ui32 == nil {
		vs := make([]uint32, len(s.vs))
		for idx, elem := range s.vs {
			ui32, err := elem.Uint32()
			if err != nil {
				return nil, err
			}
			vs[idx] = ui32
		}
		s.ui32 = vs
	}
	return s.ui32, nil
}

func (s *Slice) Uint64() ([]uint64, error) {
	if s.ui64 == nil {
		vs := make([]uint64, len(s.vs))
		for idx, elem := range s.vs {
			ui64, err := elem.Uint64()
			if err != nil {
				return nil, err
			}
			vs[idx] = ui64
		}
		s.ui64 = vs
	}
	return s.ui64, nil
}

func (s *Slice) Float32() ([]float32, error) {
	if s.f32 == nil {
		vs := make([]float32, len(s.vs))
		for idx, elem := range s.vs {
			f32, err := elem.Float32()
			if err != nil {
				return nil, err
			}
			vs[idx] = f32
		}
		s.f32 = vs
	}
	return s.f32, nil
}

func (s *Slice) Float64() ([]float64, error) {
	if s.f64 == nil {
		vs := make([]float64, len(s.vs))
		for idx, elem := range s.vs {
			f64, err := elem.Float64()
			if err != nil {
				return nil, err
			}
			vs[idx] = f64
		}
		s.f64 = vs
	}
	return s.f64, nil
}

func (s *Slice) Bool() ([]bool, error) {
	if s.b == nil {
		vs := make([]bool, len(s.vs))
		for idx, elem := range s.vs {
			b, err := elem.Bool()
			if err != nil {
				return nil, err
			}
			vs[idx] = b
		}
		s.b = vs
	}
	return s.b, nil
}

func (s *Slice) String() []string {
	if s.str == nil {
		vs := make([]string, len(s.vs))
		for idx, elem := range s.vs {
			vs[idx] = elem.String()
		}
		s.str = vs
	}
	return s.str
}

func ToSlice(elemKind reflect.Kind, vs []interface{}) (interface{}, error) {
	slice, err := NewSlice(vs)
	if err != nil {
		return nil, err
	} else {
		switch elemKind {
		case reflect.String:
			return slice.String(), nil
		case reflect.Int:
			return slice.Int()
		case reflect.Int8:
			return slice.Int8()
		case reflect.Int16:
			return slice.Int16()
		case reflect.Int32:
			return slice.Int32()
		case reflect.Int64:
			return slice.Int64()
		case reflect.Uint:
			return slice.Uint()
		case reflect.Uint8:
			return slice.Uint8()
		case reflect.Uint16:
			return slice.Uint16()
		case reflect.Uint32:
			return slice.Uint32()
		case reflect.Uint64:
			return slice.Uint64()
		case reflect.Float32:
			return slice.Float32()
		case reflect.Float64:
			return slice.Float64()
		case reflect.Bool:
			return slice.Bool()
		default:
			return nil, fmt.Errorf("convert.ToSlice: unsupported slice element type: %s", elemKind)
		}
	}
}
