package main

import (
	"fmt"
	"github.com/polinaspasova/week3/lecture8/task1/carddraw"
	"github.com/polinaspasova/week3/lecture8/task1/cardgame"

)

func main() {

	deck:=cardgame.New()
	fmt.Println(*deck)
	fmt.Print("\n\n")
    fmt.Println(carddraw.DrawAllCards(deck))

}
