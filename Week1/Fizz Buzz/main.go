package main

import "fmt"

func main() {

	var n int = 100

	//	fmt.Println("Enter N : ")
	//	fmt.Scanln(n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz ")
		} else if i%5 == 0 {
			fmt.Println("Buzz ")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz ")
		} else {
			fmt.Println(i)
		}
	}
}
