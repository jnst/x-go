package code

import (
	"strings"
)

// Pair represents key value pair.
type Pair struct {
	K string
	V int
}

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
		}
	}

	return m
}
