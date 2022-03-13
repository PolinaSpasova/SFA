package carddraw

import (
	"github.com/polinaspasova/week3/lecture8/task1/cardgame"
)

type dealer interface{
	Deal() *cardgame.Card
}

func DrawAllCards(d dealer) []cardgame.Card {
	// call the dealer's Draw() method, until you reach a nil Card
	var res []cardgame.Card
	for {
		if i:=d.Deal();i!=nil{
		res = append(res, *i)
		}else{
			break;
		}
	}
	return res
}
