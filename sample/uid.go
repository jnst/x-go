package sample

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/rs/xid"
)

func main() {
	fmt.Printf("========== XID ==========\n\n")
	printXID()
	fmt.Printf("\n========== ULID ==========\n\n")
	printULID()
}

func printXID() {
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

func printULID() {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	for i := 0; i < 5; i++ {
		fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy))
	}

	time.Sleep(1 * time.Second)
	fmt.Println("-----")

	for i := 0; i < 5; i++ {
		fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy))
	}
}
