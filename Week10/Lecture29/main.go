package main

import "fmt"

type Order struct {
	Customer string
	Amount   int
}

func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]T {
	outs := make(map[U][]T, len(col))

	for _, item := range col {
		outs[keyFn(item)] = append(outs[keyFn(item)], item)
	}

	return outs
}

func main() {

	new := []Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}

	results := GroupBy(new, func(o Order) string { return o.Customer })

	fmt.Println(results)
}
