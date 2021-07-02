package game_test

import (
	"testing"

	"github.com/jnst/x-go/game"
)

func TestCalcGem(t *testing.T) {
	t.Parallel()

	type AmountTest struct {
		in  int
		out int
	}

	tests := []AmountTest{
		{-1, 0},
		{0, 0},
		{1, 1},
		{1000, 5},
		{1001, 5},
		{2000, 7},
		{10000, 25},
		{100000, 125},
		{1000000, 600},
		{10000000, 3000},
	}

	for _, test := range tests {
		c := game.NewCalculator()
		actual := c.CalcGem(test.in)

		if actual != test.out {
			t.Errorf("CalcGem(%v) = %v; expected %v", test.in, actual, test.out)
		}
	}
}
