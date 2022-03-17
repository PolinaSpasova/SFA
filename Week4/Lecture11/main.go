package main

import (
	"fmt"
	"sync"
)

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8}
	evenCh := processEven(inputs)
	for v:=range evenCh {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	oddCh := processOdd(inputs)
	for v:=range oddCh {
		fmt.Printf("%d ", v)
	}

}

func processEven(inputs []int) chan int {
	var wg sync.WaitGroup
	response := make(chan int)
	go func() {
		for _, value := range inputs {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				if val%2 == 0 {
					response <- val
				}
			}(value)
		}
		wg.Wait()
		close(response)
	}()
	return response
}

func processOdd(inputs []int) chan int {
	var wg sync.WaitGroup
	response := make(chan int)
	go func() {
		for _, value := range inputs {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				if val%2 != 0 {
					response <- val
				}
			}(value)
		}
		wg.Wait()
		close(response)
	}()
	return response
}