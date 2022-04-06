package primes

import (
	"sync"
	"time"
)

func PrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}

func GoPrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	var wg sync.WaitGroup
	for k := 2; k < n; k++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i {
						res = append(res, k)
					}
					break
				}
			}
		}()
		wg.Wait()
	}
	return res
}
