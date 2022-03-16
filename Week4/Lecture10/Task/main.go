package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	state int
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				if cp.state == idx {
					cp.Lock()
					fmt.Print("foo")
					cp.state++
					cp.Unlock()
					return
				}
			}
		}(i)
		i++
	}

	time.Sleep(10 * time.Millisecond)
}

func (cp *ConcurrentPrinter) printBar(times int) {
	var wg sync.WaitGroup
	for i := 1; i < times; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				if cp.state == idx {
					cp.Lock()
					fmt.Print("bar")
					cp.state++
					cp.Unlock()
					return
				}
			}
		}(i)
		i++
	}

	time.Sleep(10 * time.Millisecond)
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.printFoo(times)
	cp.printBar(times)
	cp.Wait()
}
