// Copyright 2020 morgine.com. All rights reserved.

package convert_test

import (
	"github.com/morgine/convert"
	"testing"
)

func TestMustGetValue(t *testing.T) {
	var (
		need bool
		got  bool
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		str  string
		f32  float32
		f64  float64
	)
	need = true
	v, err := convert.GetValue(need)
	handleErr(err, t)
	i, err = v.Int()
	handleErr(err, t)
	i8, err = convert.Int(i).Int8()
	handleErr(err, t)
	i16, err = convert.Int8(i8).Int16()
	handleErr(err, t)
	i32, err = convert.Int16(i16).Int32()
	handleErr(err, t)
	i64, err = convert.Int32(i32).Int64()
	handleErr(err, t)
	ui, err = convert.Int64(i64).Uint()
	handleErr(err, t)
	ui8, err = convert.Uint(ui).Uint8()
	handleErr(err, t)
	ui16, err = convert.MustGetValue(ui8).Uint16()
	handleErr(err, t)
	ui32, err = convert.Uint16(ui16).Uint32()
	handleErr(err, t)
	ui64, err = convert.Uint32(ui32).Uint64()
	handleErr(err, t)
	str = convert.Uint64(ui64).String()
	handleErr(err, t)
	f32, err = convert.String(str).Float32()
	handleErr(err, t)
	f64, err = convert.Float32(f32).Float64()
	handleErr(err, t)
	got, err = convert.Float64(f64).Bool()
	handleErr(err, t)
	if got != need {
		t.Fatalf("need: %v, got: %v", need, got)
	}
}

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
