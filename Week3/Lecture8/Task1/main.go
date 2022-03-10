package main

import (
	"fmt"
	
)

func main() {

	new := cardgame.New()
	fmt.Println(*new)
	fmt.Println()
	sh := new.Shuffle()
	fmt.Println(*sh)
	fmt.Println()
	for i := 0; i < 53; i++ {
		deal := sh.Deal()
		fmt.Println(*deal)
	}
	/*
		a:=Deck{}
		fmt.Println(a)
		fmt.Println(a.Deal())
	*/
}
