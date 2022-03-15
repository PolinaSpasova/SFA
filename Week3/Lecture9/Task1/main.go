package main

import (
	"fmt"
	"log"

	"github.com/polinaspasova/Week3/Lecture9/Task1/carddrawTwo"
	"github.com/polinaspasova/Week3/Lecture9/Task1/cardgameTwo"
)

func main() {
	//deck:=cardgameTwo.New()
	deck := cardgameTwo.Deck{}
	fmt.Println(deck)
	fmt.Print("\n\n")
	//	fmt.Println(deck.Deal())
	res, err := carddrawTwo.DrawAllCards(&deck)
	if err != nil {
		
		log.Fatalf("Error %s occured ", err)
	} else {
		fmt.Println(res)
	}
}
