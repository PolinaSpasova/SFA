package main

import (
	"fmt"
)

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

/*func maxCard(cards []Card) Card {
	max := cards[0]
	for idx := range cards {
		if compareCards(cards[idx], max) == 1 {
			max = cards[idx]
		}
	}
	return max
}*/

type CardComparator func(cOne Card, cTwo Card) int

func maxCard(cards []Card, comparatorFunc CardComparator) Card {
	max:=cards[0]
	for idx:=range cards{
		if comparatorFunc(cards[idx],max)==1{
			max=cards[idx]
		}
	}
	return max
}


func main() {

	/*first := Card{6, 3}
	second := Card{5, 3}
	i := compareCards(first, second)
	switch i {
	case 1:
		fmt.Printf("%d => The first card has greater strenght", i)
	case -1:
		fmt.Printf("%d => The second card has greater strenght", i)
	case 0:
		fmt.Printf("%d => The cards have equal strenght ", i)
	}*/

	slice := []Card{{value: 5, suit: 4},
		{value: 13, suit: 2},
		{value: 12, suit: 1},
		{value: 14, suit: 4},
		{value: 3, suit: 2},
		{value: 9, suit: 3},
		{value: 2, suit: 1},
		{value: 13, suit: 4},
	}

	/*slice := make([]Card, 8)
	slice[0] = Card{5, 4}
	slice[1] = Card{13, 2}
	slice[2] = Card{12, 1}
	slice[3] = Card{6, 4}
	slice[4] = Card{3, 2}
	slice[5] = Card{9, 3}
	slice[6] = Card{13, 4}
	slice[7] = Card{13, 4}*/

	fmt.Print("The max card is ")
	fmt.Println(maxCard(slice,compareCards))
	

}
