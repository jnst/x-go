package main

import "testing"

func TestFactorial(t *testing.T) {
	type NumberTest struct {
		n      int64
		expect int64
	}

	var tests = []NumberTest{
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
		actual := Factorial(test.n)
		if actual != test.expect {
			t.Errorf("factorial(%v) = %v; expected %v", test.n, actual, test.expect)
		}
	}
}
