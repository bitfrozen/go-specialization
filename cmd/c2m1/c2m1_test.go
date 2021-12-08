package main

import (
	"math"
	"reflect"
	"testing"
)

func TestBubbleSortWithInvalidData(t *testing.T) {
	var testSlice []int = nil
	BubbleSort(testSlice)
	if testSlice != nil {
		t.Errorf("input: %v, expected: %v, got: %v", nil, nil, testSlice)
	}
}

func TestBubbleSortWithValidData(t *testing.T) {
	type testCase struct {
		input []int
		want  []int
	}

	testCases := []testCase{
		{input: []int{0}, want: []int{0}},
		{input: []int{1, 0}, want: []int{0, 1}},
		{input: []int{-1, 0}, want: []int{-1, 0}},
		{input: []int{0, 0}, want: []int{0, 0}},
		{input: []int{-1, -555}, want: []int{-555, -1}},
		{input: []int{math.MaxInt, 0, math.MinInt}, want: []int{math.MinInt, 0, math.MaxInt}},
		{input: []int{0, 1, 9, 546, 23, 12, 45, 3, 2, 1}, want: []int{0, 1, 1, 2, 3, 9, 12, 23, 45, 546}},
		{input: []int{-4565, 12, 654, 21, 56, 1, -446, 0, 4, 5}, want: []int{-4565, -446, 0, 1, 4, 5, 12, 21, 56, 654}},
	}

	for _, c := range testCases {
		testSlice := make([]int, len(c.input))
		copy(testSlice, c.input)
		BubbleSort(testSlice)
		if !reflect.DeepEqual(testSlice, c.want) {
			t.Errorf("input: %v, expected: %v, got: %v", c.input, c.want, testSlice)
		}
	}
}
