package code_test

import (
	"testing"

	"github.com/jnst/x-go/code"
)

func TestSumEachLevelDomain(t *testing.T) {
	pairs := []code.Pair{
		{"news.example.com", 100},
		{"example.com", 220},
		{"google.com", 500},
		{"example.jp", 30},
	}
	output := code.SumEachDomainHierarchy(pairs)

	expect := 6
	if len(output) != expect {
		t.Fatalf("expect: %d, output: %d", expect, len(output))
	}

	expect = 820
	if count, ok := output["com"]; !ok || count != expect {
		t.Fatalf("expect: %d, output: %d", expect, count)
	}
}
