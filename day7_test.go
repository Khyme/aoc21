package main

import (
	"strconv"
	"testing"
)

func TestDay7(t *testing.T) {
	// Input data
	ex := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	if crabs(ex) != 37 {
		t.Errorf("Example got %d; want 37", crabs(ex))
	}
	rec := readCsvFile(7)
	ints := make([]int, len(rec[0]))
	for i, val := range rec[0] {
		valInt, _ := strconv.Atoi(val)
		ints[i] = valInt
	}
	if crabs(ints) != 348664 {
		t.Errorf("Example got %d; want 348664", crabs(ints))
	}
}
