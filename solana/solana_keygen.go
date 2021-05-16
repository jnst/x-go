package solana

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func GenerateVanityAddress(prefix string) {
	start := time.Now()
	cnt := 1
	for {
		out, err := exec.Command(
			"solana-keygen",
			"new",
			"--no-outfile",
			"--no-bip39-passphrase",
			"--word-count",
			"24",
		).Output()
		if err != nil {
			panic(err)
		}

		s := string(out)
		pubkey := strings.Split(s, "\n")[2][8:]

		if strings.HasPrefix(pubkey, prefix) {
			fmt.Printf("%sdone! %f sec\n", s, time.Now().Sub(start).Seconds())
			return
		}

		if cnt%1000 == 0 {
			fmt.Printf("%d: %f sec\n", cnt, time.Now().Sub(start).Seconds())
		}
		cnt++
	}
}
