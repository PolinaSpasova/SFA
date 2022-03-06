package main

import "fmt"

type Card struct {
	value int
	suit  int
}

const (
	club = iota + 1
	diamond
	heart
	spade
)

func compareCards(cardOne Card, cardTwo Card) int {
	if (cardOne.value >= 2 && cardOne.value <= 14) && (cardOne.suit >= club && cardOne.suit <= spade) && (cardTwo.value >= 2 && cardTwo.value <= 14) && (cardTwo.suit >= club && cardTwo.suit <= spade) {
		if (cardOne.value > cardTwo.value && cardOne.suit > cardTwo.suit) || (cardOne.value < cardTwo.value && cardOne.suit > cardTwo.suit) || (cardOne.value > cardTwo.value && cardOne.suit == cardTwo.suit) || (cardOne.value == cardTwo.value && cardOne.suit > cardTwo.suit) {
			return 1
		} else if (cardOne.value > cardTwo.value && cardOne.suit < cardTwo.suit) || (cardOne.value < cardTwo.value && cardOne.suit < cardTwo.suit) || (cardOne.value < cardTwo.value && cardOne.suit == cardTwo.suit) || (cardOne.value == cardTwo.value && cardOne.suit < cardTwo.suit) {
			return -1
		} else if cardOne.value == cardTwo.value && cardOne.suit == cardTwo.suit {
			return 0
		} 
	}
	fmt.Println("Inavlid values! Card value must be between 2 and 14 , and the suit from 1 to 4! ")
			return -2
}

func main() {

	first:=Card{6,3}
	second:=Card{5,3}
	i:=compareCards(first,second)
	switch i{
	case 1:
		fmt.Printf("%d => The first card has greater strenght",i)
	case -1:
		fmt.Printf("%d => The second card has greater strenght",i)
	case 0:
		fmt.Printf("%d => The cards have equal strenght ",i)
	}

}