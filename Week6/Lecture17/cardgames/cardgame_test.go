package cardgames

import "testing"

func TestMaxCard(t *testing.T) {
	cases := []Card{{Value: 5, Suit: 4},
	{Value: 13, Suit: 2},
	{Value: 12, Suit: 1},
	{Value: 14, Suit: 4},
	{Value: 3, Suit: 2},
	{Value: 9, Suit: 3},
	{Value: 2, Suit: 1},
	{Value: 13, Suit: 4},
}
	got:=MaxCard(cases)
	expected:=Card{Value: 14,Suit:4}

	if got!=expected{
		t.Errorf("expected %v , got %v" ,expected, got)
	}
}

