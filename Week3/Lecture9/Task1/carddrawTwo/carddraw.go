package carddrawTwo

import (
	"github.com/polinaspasova/Week3/Lecture9/Task1/cardgameTwo"
)

type dealer interface {
	Deal() (*cardgameTwo.Card, error)
	Done() bool
}

func DrawAllCards(d dealer) ([]cardgameTwo.Card, error) {
	// call the dealer's Draw() method, until you reach a nil Card
	var res []cardgameTwo.Card
	for {
		card, err := d.Deal()
		if err != nil {
			if d.Done() {
				return res, nil
			}else {
				return nil,err
			}
		} else {
			res = append(res, *card)
		}
	}
	//return res, nil
}
