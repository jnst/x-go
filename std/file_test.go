package std_test

import (
	"github.com/jnst/x-go/std"
	"testing"
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
