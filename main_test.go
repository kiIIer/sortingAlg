package main

import (
	"fmt"
	"testing"
)

type testRandStruct struct {
	calledCount int
	values      []int
}

func (r *testRandStruct) Intn(max int) int {
	r.calledCount++

	return r.values[r.calledCount-1]
}

func TestMakeSlice(t *testing.T) {

	var testCase = []struct {
		name     string
		leng     int
		expected []int
	}{
		{"1", 6, []int{1, 2, 3, 6, 5, 4}},
		{"2", 10, []int{1, 2, 3, 6, 5, 4, 1, 2, 3, 6}},
	}

	for _, tt := range testCase {
		testName := fmt.Sprint(tt.name)
		t.Run(testName, func(t *testing.T) {
			randValues := make([]int, len(tt.expected))
			for i, v := range tt.expected {
				randValues[i] = v - 1
			}
			testRand := &testRandStruct{
				values: randValues,
			}

			actual := makeSlice(testRand, tt.leng, 100)

			for i := range tt.expected {
				if actual[i] != tt.expected[i] {
					t.Errorf("Houston we have a problem!!! Got %d, want %d", actual, tt.expected)
					break
				}
			}
		})
	}
}

func TestSort(t *testing.T) {
	var testCase = []struct {
		name     string
		s        []int
		expected []int
	}{
		{
			"1", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"2", []int{1, 3, 2, 4, 5}, []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range testCase {
		testName := fmt.Sprint(tt.name)
		t.Run(testName, func(t *testing.T) {
			sort(tt.s)
			actual := tt.s
			for i := range tt.s {
				if tt.s[i] != actual[i] {
					t.Errorf("Houston we have a problem!!! Got %d, want %d", tt.s, tt.expected)
				}
			}
		})
	}
}
