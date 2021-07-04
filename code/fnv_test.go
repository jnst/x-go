package code_test

import (
	"testing"

	"github.com/jnst/x-go/code"
)

func TestHash(t *testing.T) {
	t.Parallel()

	type TextTest struct {
		in  string
		out uint32
	}

	tests := []TextTest{
		{in: "", out: 2166136261},
		{in: "a", out: 84696446},
		{in: "ab", out: 1886858552},
		{in: "abc", out: 1134309195},
	}

	for _, test := range tests {
		got := code.Hash(test.in)
		if got != test.out {
			t.Errorf("Hash(%s) = %d want %d", test.in, got, test.out)
		}
	}
}
