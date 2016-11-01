package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

// RunSay is multi thread sample
func RunSay() {
	go say("world")
	say("hello")
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

// RunSum is channel sample
func RunSum() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// Fail occurs deadlock
func Fail() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// RunFibonacci is fibonacci sample
func RunFibonacci() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// ShowChannelCapacity shows channel capacity
func ShowChannelCapacity() {
	c1 := make(chan int)
	c2 := make(chan int, 10)
	fmt.Println("make(chan int)     ->", cap(c1))
	fmt.Println("make(chan int, 10) ->", cap(c2))
}

// RunTimeout is timeout sample
func RunTimeout() {
	println("start..")
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	println(<-o)
}

// ShowCPUCoreCount shows cpu core count
func ShowCPUCoreCount() {
	println(runtime.NumCPU)
}
