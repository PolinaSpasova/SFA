package main

import "fmt"

type CardSuit = int

const (
	club CardSuit = iota
	diamond
	heart
	spade
	
)

const (
	two = 2 + iota
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	j
	q
	k
	a
)

func compareCards(card1Val int, card1Suit int, card2Val int, card2Suit int) int {
	if (card1Val>=two&&card1Val<=a) && (card1Suit>=club&&card1Suit<=spade) && (card2Val>=two&&card2Val<=a) && (card2Suit>=club&&card2Suit<=spade){
		if (card1Val>card2Val && card1Suit>card2Suit) || (card1Val<card2Val && card1Suit>card2Suit) || (card1Val>card2Val && card1Suit==card2Suit)|| (card1Val==card2Val && card1Suit>card2Suit){
			return 1
		}else if (card1Val>card2Val && card1Suit<card2Suit) || (card1Val<card2Val && card1Suit<card2Suit)|| (card1Val<card2Val && card1Suit==card2Suit)|| (card1Val==card2Val && card1Suit<card2Suit){
			return -1
		}else if (card1Val==card2Val && card1Suit==card2Suit)  {
			return 0
		}
	}
	fmt.Println("Inavlid values!")
	return -2
	
}

func main() {

fmt.Println(compareCards(14,2,14,1))
}
