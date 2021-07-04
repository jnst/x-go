package atcoder

import "fmt"

func Welcome() {
	var (
		a, b, c int
		s string
	)

	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d %d", &b, &c)
	_, _ = fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}
