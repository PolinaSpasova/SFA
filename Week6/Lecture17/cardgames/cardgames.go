package cardgames

import "fmt"

type Card struct {
	Value int
	Suit  int
}

const (
	club = iota + 1
	diamond
	heart
	spade
)

func compareCards(cardOne Card, cardTwo Card) int {
	if (cardOne.Value >= 2 && cardOne.Value <= 14) && (cardOne.Suit >= club && cardOne.Suit <= spade) && (cardTwo.Value >= 2 && cardTwo.Value <= 14) && (cardTwo.Suit >= club && cardTwo.Suit <= spade) {
		if (cardOne.Value > cardTwo.Value && cardOne.Suit > cardTwo.Suit) || (cardOne.Value < cardTwo.Value && cardOne.Suit > cardTwo.Suit) || (cardOne.Value > cardTwo.Value && cardOne.Suit == cardTwo.Suit) || (cardOne.Value == cardTwo.Value && cardOne.Suit > cardTwo.Suit) {
			return 1
		} else if (cardOne.Value > cardTwo.Value && cardOne.Suit < cardTwo.Suit) || (cardOne.Value < cardTwo.Value && cardOne.Suit < cardTwo.Suit) || (cardOne.Value < cardTwo.Value && cardOne.Suit == cardTwo.Suit) || (cardOne.Value == cardTwo.Value && cardOne.Suit < cardTwo.Suit) {
			return -1
		} else if cardOne.Value == cardTwo.Value && cardOne.Suit == cardTwo.Suit {
			return 0
		}
	}
	fmt.Println("Inavlid values! Card value must be between 2 and 14 , and the suit from 1 to 4! ")
	return -2
}

func MaxCard(cards []Card) Card {
	max := cards[0]
	for idx := range cards {
		if compareCards(cards[idx], max) == 1 {
			max = cards[idx]
		}
	}
	return max 
}
