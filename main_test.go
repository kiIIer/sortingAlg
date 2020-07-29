package main

import (
	"fmt"
	"testing"
)

var tests1 = []struct {
	leng, len int
}{
	{1, 1},
	{10, 10},
}

func TestMakeSlice(t *testing.T) {
	for _, tt := range tests1 {
		testname := fmt.Sprintf("%d, %d", tt.leng, tt.len)
		t.Run(testname, func(t *testing.T) {
			ans := len(makeSlice(tt.leng, 100))
			if ans != tt.len {
				t.Errorf("Houston we have a problem!!! Got %d, want %d", ans, tt.len)
			}
		})
	}
}

var tests2 = []struct {
	leng int
	s    []int
	ans  []int
}{
	{
		10, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		5, []int{1, 3, 2, 4, 5}, []int{1, 2, 3, 4, 5},
	},
}

func TestSort(t *testing.T) {
	for _, tt := range tests2 {
		testname := fmt.Sprintf("%d, %d, %d", tt.leng, tt.s, tt.ans)
		t.Run(testname, func(t *testing.T) {
			sort(tt.leng, tt.s)
			for i := 0; i < tt.leng; i++ {
				if tt.s[i] != tt.ans[i] {
					t.Errorf("Houston we have a problem!!! Got %d, want %d", tt.s, tt.ans)
				}
			}
		})
	}
}
