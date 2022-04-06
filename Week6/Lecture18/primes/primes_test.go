package primes

import (
	"testing"
	"time"
)

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesAndSleep(100, 0*time.Millisecond)
	}
}

func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesAndSleep(100, 5*time.Millisecond)
	}
}
func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesAndSleep(100, 10*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPrimesAndSleep(100, 0*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPrimesAndSleep(100, 5*time.Millisecond)
	}
}
func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPrimesAndSleep(100, 10*time.Millisecond)
	}
}
