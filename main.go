package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type rnd interface {
	Intn(max int) int
}

func makeSlice(rand rnd, leng, maxn int) []int {
	s := make([]int, leng)
	for i := 0; i < leng; i++ {
		s[i] = rand.Intn(maxn) + 1
	}
	return s
}

func sort(s []int) {
	for {
		changed := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}

func main() {
	lengOfSlice := flag.Int("leng", 10, "length of slice")
	maxNumber := flag.Int("maxn", 100, "max number in slice")
	flag.Parse()
	if *lengOfSlice < 1 || *maxNumber < 0 {
		fmt.Println("You weren't supposed to do that")
		return
	}
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	s := makeSlice(rand, *lengOfSlice, *maxNumber)

	fmt.Printf("%+v \n", s)
	fmt.Println("Sorting...")

	sort(s)

	fmt.Printf("%+v \n", s)
	fmt.Println("Done!")
}
