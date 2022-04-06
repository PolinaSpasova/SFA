package main

import (
	"fmt"
	"time"
	"github.com/polinaspasova/Week6/Lecture18/primes"
)

func main() {
	fmt.Println(primes.PrimesAndSleep(10, 10*time.Nanosecond))
	fmt.Println(primes.GoPrimesAndSleep(10, 10*time.Nanosecond))

}

