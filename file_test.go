package main

import "testing"

func TestList(t *testing.T) {
	expect := "file.go"
	for _, v := range List() {
		if v == expect {
			return
		}
	}
	t.Errorf("%s not found.", expect)
}
