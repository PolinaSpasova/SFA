package main

import (
	"fmt"
)

func main() {
	ch := fibonacci(10)
	for v := range ch {
		fmt.Println(v)
	}

}

func fibonacci(n int) chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			ch <- a
			a, b = b, a+b
		}
		close(ch)
	}()
	return ch
}
