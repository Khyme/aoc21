package main

func lantern(start []int, days int) int {

	startmap := make(map[int]int)
	for i := 0; i < 9; i++ {
		startmap[i] = 0
	}
	for _, val := range start {
		startmap[val] = startmap[val] + 1
	}
	tot := 0
	res, _ := lanternRec2(startmap, days)
	for i := 0; i < 9; i++ {
		tot = tot + res[i]
	}
	return tot
}

func lanternRec(start []int, days int) ([]int, int) {
	if days == 0 {
		return start, 0
	}

	end := make([]int, len(start))
	copy(end, start)
	for i, val := range start {
		if val == 0 {
			end[i] = 6
			end = append(end, 8)
		} else {
			end[i] = val - 1
		}
	}

	return lanternRec(end, days-1)
}

func lanternRec2(startmap map[int]int, days int) (map[int]int, int) {

	if days == 0 {
		return startmap, 0
	}

	newmap := make(map[int]int)

	for i := 0; i < 9; i++ {
		if i == 0 {
			newmap[8] = startmap[i]
			newmap[6] = startmap[i]
		} else {
			newmap[i-1] = newmap[i-1] + startmap[i]
		}
	}

	return lanternRec2(newmap, days-1)
}
