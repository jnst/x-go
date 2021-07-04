package std_test

import (
	"testing"

	"github.com/jnst/x-go/std"
)

func TestList(t *testing.T) {
	expect := "file.go"
	for _, v := range std.List() {
		if v == expect {
			return
		}
	}

	t.Errorf("%s not found.", expect)
}
