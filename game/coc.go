package game

import "math"

var (
	defaultResourceRanges = []int{100, 1000, 10000, 100000, 1000000, 10000000}
	defaultRequiredGems   = []int{1, 5, 25, 125, 600, 3000}
)

type Calculator struct {
	resourceRanges []int
	requiredGems   []int
}

func NewCalculator() *Calculator {
	return &Calculator{
		resourceRanges: defaultResourceRanges,
		requiredGems:   defaultRequiredGems,
	}
}

// CalcGem returns gem price.
// Clash of Clans specification for purchase resource by gem.
func (c *Calculator) CalcGem(value int) int {
	if value <= 0 {
		return 0
	}

	if value <= c.resourceRanges[0] {
		return c.requiredGems[0]
	}

	for i, v := range c.resourceRanges {
		if value <= v {
			return c.calc(i, value)
		}
	}

	return c.calc(len(c.resourceRanges)-1, value)
}

func (c *Calculator) calc(i, v int) int {
	x1 := float64(c.requiredGems[i])
	x2 := float64(c.requiredGems[i-1])
	y1 := float64(c.resourceRanges[i])
	y2 := float64(c.resourceRanges[i-1])
	result := (x1-x2)/(y1-y2)*(float64(v)-y2) + x2

	return int(math.Floor(result + .5))
}
