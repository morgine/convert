// Copyright 2020 morgine.com. All rights reserved.

package convert_test

import (
	"fmt"
	"github.com/morgine/convert"
)

func ExampleMultiples() {
	var (
		any convert.Value // all convert implemented the interface
		i32 int32
		i64 int64
		str string
		err error
	)

	any = convert.MustGetValue(1987) // automatic conversion type

	i32, _ = any.Int32()

	str = convert.Int32(i32).String()

	i64, _ = convert.String(str).Int64()

	_, err = convert.Int64(i64).Int8()

	fmt.Println(err == convert.ErrOutOfInt8Range)

	// Output:
	// true
}

func ExampleFloatPrecision() {
	var (
		f1, f2 float64
	)
	f1 = 999999999999999999
	f2 = f1 + 10
	fmt.Println(f1 == f2)

	// Output:
	// true
}

func ExampleNewSlice() {
	var (
		ints   = []int{1, 2, 3}
		strs   []string
		floats []float64
	)
	vars := convert.ReflectToSlice(ints)
	slice, err := convert.NewSlice(vars)
	if err != nil {
		panic(err)
	}
	strs = slice.String()
	fmt.Println(strs)

	floats, err = slice.Float64()
	if err != nil {
		panic(err)
	}
	fmt.Println(floats)

	// Output:
	// [1 2 3]
	// [1 2 3]
}
