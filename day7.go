package main

import (
	"fmt"
	"math"
	"sort"
)

func crabs(start []int) int {
	sort.Ints(start)
	tot := len(start)
	sum := 0
	med := 0
	for i := 0; i < tot; i++ {
		sum += start[i]
		if i == tot/2 {
			med = start[i]
		}
	}
	medF := float64(med)
	meanF := math.Round(float64(sum) / float64(tot))
	min1 := math.MaxInt
	min2 := math.MaxInt
	var j1, j2 float64
	for j := math.Min(medF, meanF); j <= math.Max(medF, meanF); j++ {
		val1 := fuel(start, int(j))
		if val1 < min1 {
			j1 = j
			min1 = val1
		}
		val2 := fuel2(start, int(j))
		if val2 < min2 {
			j2 = j
			min2 = val2
		}
	}
	fmt.Printf("%v/%v %v/%v\n", j1, min1, j2, min2)
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
