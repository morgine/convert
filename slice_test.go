// Copyright 2020 morgine.com. All rights reserved.

package convert_test

import (
	"github.com/morgine/convert"
	"testing"
)

// 567 ns/op
func BenchmarkNewSlice(b *testing.B) {
	var vars = convert.ReflectToSlice([]int{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		slice, err := convert.NewSlice(vars)
		if err != nil {
			b.Fatal(err)
		}
		slice.String()
	}
}

// 359 ns/op
func BenchmarkReflectToSlice(b *testing.B) {
	var vars = []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		slice := convert.ReflectToSlice(vars)
		if got, need := len(slice), len(vars); got != need {
			b.Fatalf("need: %d, got: %d", need, got)
		}
	}
}
