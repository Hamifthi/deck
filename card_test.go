package deck

import (
	"fmt"
	"testing"
)

func TestCardString(t *testing.T) {
	var tests = []struct {
		card Card
		want string
	}{
		{Card{Rank: Ace, Suit: Heart}, "Ace of Hearts"},
		{Card{Rank: Two, Suit: Diamond}, "Two of Diamonds"},
		{Card{Rank: Eight, Suit: Club}, "Eight of Clubs"},
		{Card{Rank: King, Suit: Spade}, "King of Spades"},
		{Card{Suit: Joker}, "Joker"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.card)
		t.Run(testname, func(t *testing.T) {
			ans := tt.card.String()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 52 {
		t.Errorf("Want %d cards get %d", 52, len(cards))
	}
}

func TestLess(t *testing.T) {
	var tests = []struct {
		i, j int
		want bool
	}{
		{3, 6, true},
		{13, 27, true},
		{40, 30, false},
	}
	cards := NewDeck()
	less := Less(cards)

	for _, tt := range tests {
		testname := fmt.Sprintf("pair of cards: %s, %s",
			cards[tt.i].String(), cards[tt.j].String())
		t.Run(testname, func(t *testing.T) {
			ans := less(tt.i, tt.j)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestDefaultSort(t *testing.T) {
	cards := NewDeck(DefaultSort)
	firstCard := Card{Suit: Spade, Rank: Ace}
	lastCard := Card{Suit: Heart, Rank: King}
	if cards[0] != firstCard && cards[len(cards)-1] != lastCard {
		t.Errorf("Sort function doesn't work correctly")
	}
}

func TestShuffleDeck(t *testing.T) {
	cards := NewDeck(ShuffleDeck())
	firstCard := Card{Suit: Spade, Rank: Ace}
	lastCard := Card{Suit: Heart, Rank: King}
	if cards[0] == firstCard || cards[len(cards)-1] == lastCard {
		t.Errorf("Shuffle function doesn't work correctly")
	}
}

func TestAddJoker(t *testing.T) {
	cards := NewDeck(AddJoker(5))
	lastCard := Card{Suit: Joker, Rank: Five}
	if len(cards) == 52+5 && cards[len(cards)-1] == lastCard {
		t.Errorf("Add Joker function doesn't work correclty")
	}
}

func TestFilter(t *testing.T) {
	cards := NewDeck(Filter(map[int]struct{}{2: struct{}{}, 3: struct{}{}}))
	for _, card := range cards {
		if int(card.Rank) == 2 || int(card.Rank) == 3 {
			t.Errorf("Got rank %d from cards", card.Rank)
		}
	}
}

func TestMultipleDeck(t *testing.T) {
	cards := NewDeck(MultipleDeck(3))
	if len(cards) != 52*3 {
		t.Errorf("Expected %d but got %d", 52*3, len(cards))
	}
}
