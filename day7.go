package main

import (
	"fmt"
	"math"
	"sort"
)

func crabs(start []int) int {
	sort.Ints(start)
	tot := len(start)
	min1 := math.MaxInt
	min2 := math.MaxInt

	for j := start[0]; j <= start[tot-1]; j++ {
		val1 := fuel(start, j)
		if val1 < min1 {
			min1 = val1
		}
		val2 := fuel2(start, j)
		if val2 < min2 {
			min2 = val2
		}
	}
	fmt.Printf("%v %v\n", min1, min2)
	return min1
}

func fuel(start []int, mid int) int {
	fuel := 0
	for _, v := range start {
		fuel += int(math.Abs(float64(v - mid)))
	}
	return fuel
}

func fuel2(start []int, mid int) int {
	fuel := 0
	for _, v := range start {
		t := int(math.Abs(float64(v - mid)))
		fuel += (t * (t + 1)) / 2
	}
	return fuel
}
