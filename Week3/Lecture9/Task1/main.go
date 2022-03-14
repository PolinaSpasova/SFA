package main

import (
	"fmt"

	"github.com/polinaspasova/Week3/Lecture9/Task1/cardgameOne"
)

func main() {
//	err:=errors.New("Missing")
	d:=cardgameOne.New()
	fmt.Println(*d)
	fmt.Print("\n\n")
	fmt.Println(d.Deal())
 //   carddrawOne.DrawAllCards(d)
}
