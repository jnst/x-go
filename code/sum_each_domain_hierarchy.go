package code

import (
	"strings"
)

// Pair represents key value pair.
type Pair struct {
	K string
	V int
}

// input:
// [
//   ("news.example.com", 100),
//   ("example.com", 220),
//   ("google.com", 500),
//   ("example.jp", 30)
// ]
// output:
// {
//   "com": 820,
//   "example.com": 320,
//   "news.example.com": 100,
//   "google.com": 500,
//   "jp": 30,
//   "example.jp": 30,
// }
func SumEachDomainHierarchy(pairs []Pair) map[string]int {
	m := map[string]int{}

	for _, pair := range pairs {
		levelDomains := strings.Split(pair.K, ".")
		size := len(levelDomains)
		for i := range levelDomains {
			domain := levelDomains[size-i-1:]
			target := strings.Join(domain, ".")
			count, ok := m[target]
			if ok {
				m[target] = count + pair.V
			} else {
				m[target] = pair.V
			}

			//fmt.Printf("%v => %s: %d\n", levelDomains, target, m[target])
		}
	}

	return m
}
