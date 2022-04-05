package main

import (
	"fmt"

	"github.com/polinaspasova/Week6/Lecture17/cardgames"
)

func main() {

	slice := []cardgames.Card{{Value: 5, Suit: 4},
		{Value: 13, Suit: 2},
		{Value: 12, Suit: 1},
		{Value: 14, Suit: 4},
		{Value: 3, Suit: 2},
		{Value: 9, Suit: 3},
		{Value: 2, Suit: 1},
		{Value: 13, Suit: 4},
	}

	fmt.Print("The max card is ")
	fmt.Println(cardgames.MaxCard(slice))


}
