//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].Suit == cards[j].Suit {
			return cards[i].Rank < cards[j].Rank
		} else {
			return cards[i].Suit < cards[j].Suit
		}
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func ShuffleDeck() func(cards []Card) []Card {
	return func(cards []Card) []Card {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
		return cards
	}
}

func AddJoker(number int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < number; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func Filter(ranks map[int]struct{}) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var newCards []Card
		for _, card := range cards {
			if _, ok := ranks[int(card.Rank)]; ok == false {
				newCards = append(newCards, card)
			}
		}
		return newCards
	}
}

func MultipleDeck(number int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var returnDeck []Card
		for i := 0; i < number; i++ {
			returnDeck = append(returnDeck, cards...)
		}
		return returnDeck
	}
}

func NewDeck(options ...func([]Card) []Card) []Card {
	var deckOfCards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deckOfCards = append(deckOfCards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, option := range options {
		deckOfCards = option(deckOfCards)
	}
	return deckOfCards
}
