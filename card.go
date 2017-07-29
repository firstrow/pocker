package pocker

import (
	"sort"
)

type Card struct {
	rank int
	suit int
}

// Cards
const (
	Two = iota + 2
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
	Ace
)

// Suits
const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

var (
	ranks = map[string]int{
		"2": Two,
		"3": Three,
		"4": Four,
		"5": Five,
		"6": Six,
		"7": Seven,
		"8": Eight,
		"9": Nine,
		"T": Ten,
		"J": Jack,
		"Q": Queen,
		"K": King,
		"A": Ace,
	}

	suits = map[string]int{
		"C": Clubs,
		"D": Diamonds,
		"H": Hearts,
		"S": Spades,
	}
)

// NewCardFromCode creates new card instance from string short code like: 2C, TS, KH
func NewCard(code string) Card {
	rank := ranks[string(code[0])]
	suit := suits[string(code[1])]
	return Card{rank, suit}
}

func (c Card) Rank() int {
	return c.rank
}

func (c Card) Suit() int {
	return c.suit
}

// byCardDesc impements sort interface
type byRankDesc []Card

func (a byRankDesc) Len() int           { return len(a) }
func (a byRankDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRankDesc) Less(i, j int) bool { return a[i].Rank() > a[j].Rank() }

// Sort cards by rank desc and returns new sorted array
func Sort(cards []Card) {
	sort.Sort(byRankDesc(cards))
}
