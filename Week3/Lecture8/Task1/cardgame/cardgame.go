package cardgame

import (
	"math/rand"
	"time"
)

type Card struct {
	value int
	suit  int
}

type Deck struct {
	deck       [52]Card
	cardsCount int
}

func New() *Deck {
	cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	suits := []int{1, 2, 3, 4}

	var newDeck Deck

	for i := range suits {
		for j := range cards {
			newDeck.deck[newDeck.cardsCount].value = cards[j]
			newDeck.deck[newDeck.cardsCount].suit = suits[i]
			newDeck.cardsCount++
		}
	}

	return &newDeck

}

func (d *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixMilli())
	for i := range d.deck {
		r := rand.Intn(i + 1)
		if i != r {
			d.deck[r], d.deck[i] = d.deck[i], d.deck[r]
		}
	}
	return d
}

func (d *Deck) Deal() *Card {
	if d.cardsCount == 0 {
		return nil
	} else {
		res := d.deck[0]
		for i := 0; i < len(d.deck)-1; i++ {
			d.deck[i] = d.deck[1+i]
		}
		d.cardsCount--
		return &res
	}
}
