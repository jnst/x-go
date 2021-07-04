package std

import "fmt"

type I interface {
	M() string
}

type T struct {
	name string
	fmt.Stringer
}

func (t T) M() string {
	return t.name
}

func (t T) String() string {
	return fmt.Sprintf("hello %s", t.name)
}

func Hello(i I) {
	fmt.Printf("hello %s\n", i.M())
}

/*func main() {
	Hello(T{name: "world"})
	fmt.Println(T{name: "world"}.String())
}*/
