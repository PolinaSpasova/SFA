package carddrawOne

import "github.com/polinaspasova/Week3/Lecture9/Task1/cardgameOne"

type dealer interface {
	Deal() *cardgameOne.Card
	Done() bool
}

func DrawAllCards(d dealer) ([]cardgameOne.Card, error) {
	// call the dealer's Draw() method, until you reach a nil Card
	var res []cardgameOne.Card
	_, err := DrawAllCards(d)
	if err != nil {
		for {
			if i := d.Deal(); i != nil {
				res = append(res, *i)
			} else {
				break
			}
		}
		return res, nil
	} else {
		if d.Done() {
			return res, nil
		} else {
			return nil, err
		}
	}

	//return nil, err
}
