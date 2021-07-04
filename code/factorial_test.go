package code_test

import (
	"testing"

	"github.com/jnst/x-go/code"
)

func TestFactorial(t *testing.T) {
	t.Parallel()

	type NumberTest struct {
		n      int64
		expect int64
	}

	tests := []NumberTest{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
		{11, 39916800},
		{12, 479001600},
		{13, 6227020800},
		{14, 87178291200},
		{15, 1307674368000},
	}

	for _, test := range tests {
		actual := code.Factorial(test.n)
		if actual != test.expect {
			t.Errorf("Factorial(%v) = %v; expected %v", test.n, actual, test.expect)
		}
	}
}
