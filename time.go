package main

import (
	"fmt"
	"time"
)

// Print prints some time values.
func Print() {
	now := time.Now()
	fmt.Println("now:", now)
	fmt.Println("unix:", now.Unix())
	fmt.Println("nano:", now.UnixNano())
}
