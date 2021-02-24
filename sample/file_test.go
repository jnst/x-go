package sample_test

import (
	"testing"

	"github.com/jnst/x-go/sample"
)

func TestList(t *testing.T) {
	expect := "file.go"
	for _, v := range sample.List() {
		if v == expect {
			return
		}
	}
	t.Errorf("%s not found.", expect)
}
