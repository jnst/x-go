package main

import "math"

var (
	ranges = []int{100, 1000, 10000, 100000, 1000000, 10000000}
	gems   = []int{1, 5, 25, 125, 600, 3000}
)

// CalcGem returns gem price.
// Clash of Clans specification for purchase resource by gem.
func CalcGem(value int) int {
	if value <= 0 {
		return 0
	}

	if value <= ranges[0] {
		return gems[0]
	}

	for i, v := range ranges {
		if value <= v {
			return calc(i, value)
		}
	}

	return calc(len(ranges)-1, value)
}

func calc(i, v int) int {
	x1 := float64(gems[i])
	x2 := float64(gems[i-1])
	y1 := float64(ranges[i])
	y2 := float64(ranges[i-1])
	result := (x1-x2)/(y1-y2)*(float64(v)-y2) + x2
	return int(math.Floor(result + .5))
}
