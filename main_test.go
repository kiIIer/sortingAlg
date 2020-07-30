package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	var testCase = []struct {
		name           string
		leng, expected int
	}{
		{"1", 1, 1},
		{"2", 10, 10},
	}

	for _, tt := range testCase {
		testName := fmt.Sprint(tt.name)
		t.Run(testName, func(t *testing.T) {
			actual := len(makeSlice(tt.leng, 100))
			if actual != tt.expected {
				t.Errorf("Houston we have a problem!!! Got %d, want %d", actual, tt.expected)
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

func TestRandNumb(t *testing.T) {
	var testCase = []struct {
		name     string
		max      int
		expected int
	}{
		{
			"1", 10, 2,
		},
		{
			"2", 100, 88,
		},
	}
	rand.Seed(1)
	for _, tt := range testCase {
		testName := fmt.Sprint(tt.name)
		t.Run(testName, func(t *testing.T) {
			actual := randNumb(tt.max)
			if actual != tt.expected {
				t.Errorf("Houston we have a problem!!! Got %d, want %d", actual, tt.expected)
			}
		})
	}
}
