package sandbox

import (
	"strings"
	"testing"
)

var (
	save   bool
	small  = 100
	medium = 1_000_000
	large  = 100_000_000
)

func Benchmark_JoinASmallSliceAndCheckContains(b *testing.B) {
	var result bool
	test := make([]string, small)
	for i := 0; i < small; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = joinContains(test, "bar")
	}
	save = result
}
func Benchmark_LoopOverASmallSliceAndCheckEquality(b *testing.B) {
	var result bool
	test := make([]string, small)
	for i := 0; i < small; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = sliceLoop(test, "bar")
	}
	save = result
}

func Benchmark_LoopOverASmallSliceAndCheckContains(b *testing.B) {
	var result bool
	test := make([]string, small)
	for i := 0; i < small; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = sliceContains(test, "bar")
	}
	save = result
}

func Benchmark_JoinAMediumSliceAndCheckContains(b *testing.B) {
	var result bool
	test := make([]string, medium)
	for i := 0; i < medium; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = joinContains(test, "bar")
	}
	save = result
}

func Benchmark_LoopOverAMediumSliceAndCheckEquality(b *testing.B) {
	var result bool
	test := make([]string, medium)
	for i := 0; i < medium; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = sliceLoop(test, "bar")
	}
	save = result
}

func Benchmark_LoopOverAMediumSliceAndCheckContains(b *testing.B) {
	var result bool
	test := make([]string, medium)
	for i := 0; i < medium; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = sliceContains(test, "bar")
	}
	save = result
}

func Benchmark_JoinALargeSliceAndCheckContains(b *testing.B) {
	var result bool
	test := make([]string, large)
	for i := 0; i < large; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = joinContains(test, "bar")
	}
	save = result
}

func Benchmark_LoopOverALargeSliceAndCheckEquality(b *testing.B) {
	var result bool
	test := make([]string, large)
	for i := 0; i < large; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = sliceLoop(test, "bar")
	}
	save = result
}

func Benchmark_LoopOverALargeSliceAndCheckContains(b *testing.B) {
	var result bool
	test := make([]string, large)
	for i := 0; i < large; i++ {
		test[i] = "foo"
	}
	for n := 0; n < b.N; n++ {
		result = sliceContains(test, "bar")
	}
	save = result
}

func joinContains(arr []string, s string) bool {
	return strings.Contains(strings.Join(arr, " "), s)
}

func sliceLoop(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func sliceContains(arr []string, s string) bool {
	for _, v := range arr {
		if strings.Contains(v, s) {
			return true
		}
	}
	return false
}
