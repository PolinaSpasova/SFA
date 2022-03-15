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
	deck       []Card
}

func New() *Deck {
	cards := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	suits := []int{1, 2, 3, 4} 
	var newDeck Deck
	var newCard Card

	for i := range suits {
		for j := range cards {
			newCard.value=cards[j]
			newCard.suit=suits[i]
			newDeck.deck=append(newDeck.deck,newCard )
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
	if len(d.deck)>0{
		res:=d.deck[0]
		d.deck=d.deck[1:]
		return &res
	}
	return nil 
}