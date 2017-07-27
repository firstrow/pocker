package pocker

import (
	"sort"
)

type Card struct {
	rank int
	suit string
}

var ranks = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var suits = map[string]string{
	"C": "clubs",
	"D": "diamonds",
	"H": "hearts",
	"S": "spades",
}

// NewCardFromCode creates new card instance from string short code like: 2C, TS, KH
func NewCard(code string) Card {
	rank := ranks[string(code[0])]
	suit := suits[string(code[1])]
	return Card{rank, suit}
}

func (c Card) Rank() int {
	return c.rank
}

func (c Card) Suit() string {
	return c.suit
}

// byCard impements sort interface
type byRank []Card

func (a byRank) Len() int           { return len(a) }
func (a byRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRank) Less(i, j int) bool { return a[i].Rank() < a[j].Rank() }

// Sort cards by rank and returns new sorted array
func Sort(cards []Card) {
	sort.Sort(byRank(cards))
}
