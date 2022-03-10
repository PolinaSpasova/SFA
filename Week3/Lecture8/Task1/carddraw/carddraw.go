package carddraw

import "github.com/PolinaSpasova/SFA/Week3/Lecture7/Task1/carddraw"

type dealer interface{
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	// call the dealer's Draw() method, until you reach a nil Card
	if dealer.Deal()!=nil{
		return dealer.Deal()
	}
	return nil
}
