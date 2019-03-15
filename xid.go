package main

import (
	"fmt"
	"github.com/rs/xid"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		guid := xid.New()
		fmt.Printf("%s: %d\n", guid.String(), guid.Counter())
	}

	time.Sleep(1 * time.Second)
	fmt.Println("-----")

	for i := 0; i < 5; i++ {
		guid := xid.New()
		fmt.Printf("%s: %d\n", guid.String(), guid.Counter())
	}
}
