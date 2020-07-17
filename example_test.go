// Copyright 2020 morgine.com. All rights reserved.

package convert_test

import (
	"fmt"
	"github.com/morgine/convert"
)

func ExampleMultiples() {
	var err error
	var value convert.Value

	value, err = convert.GetValue(1987) // automatic convert types
	if err != nil {
		panic(err)
	}
	fmt.Println(value.String() == "1987")

	value = convert.Bool(false)
	i, _ := value.Int()
	fmt.Println(i == 0)

	value = convert.String("1")
	i16, _ := value.Int16()
	fmt.Println(i16 == int16(1))

	value = convert.Int64(1987)
	_, err = value.Int8()
	fmt.Println(err == convert.ErrOutOfInt8Range)

	// Output:
	// true
	// true
	// true
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
