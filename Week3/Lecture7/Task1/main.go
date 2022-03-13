package main

import (
	"fmt"
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
	// var res Card
	// if d.cardsCount == 0 {
	// 	return nil
	// } else {
	// 	res = d.deck[0]
	// 	for i := 0; i < len(d.deck)-1; i++ {
	// 		d.deck[i] = d.deck[1+i]
	// 	}
	// 	d.cardsCount--
	// }
	// return &res
	for i:=len(d.deck);i>0;{
		res:=d.deck[0]
		d.deck=d.deck[1:]
		return &res
	}
	return nil
}

func main() {

	new := New()
	fmt.Println(*new)
	fmt.Println()
	sh := new.Shuffle()
	fmt.Println(*sh)
	fmt.Println(len(new.deck))
	fmt.Println(len(sh.deck))
	for i := 0; i < 52; i++ {
		fmt.Print(*new.Deal())
	}
	
	
	/*
		a:=Deck{}
		fmt.Println(a)
		fmt.Println(a.Deal())
	*/
}
