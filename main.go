package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func randNumb(max int) int {
	return rand.Intn(max) + 1
}

func makeSlice(leng, maxn int) []int {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, leng)
	for i := 0; i < leng; i++ {
		s[i] = randNumb(maxn)
	}
	return s
}

func sort(leng int, s []int) {
	for {
		changed := false
		for i := 0; i < leng-1; i++ {
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
	if *lengOfSlice < 0 || *maxNumber < 0 {
		fmt.Println("You weren't supposed to do that")
		return
	}

	s := makeSlice(*lengOfSlice, *maxNumber)

	fmt.Printf("%+v \n", s)
	fmt.Println("Sorting...")

	sort(*lengOfSlice, s)

	fmt.Printf("%+v \n", s)
	fmt.Println("Done!")
}
