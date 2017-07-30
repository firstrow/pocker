package pocker

import ()

type Hand int
type Evaluator func([]Card) Hand

const (
	NoHand = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

// pairsHash is a hash table which keys is card rank
// and value is number of occurences in hand.
// eg: map[rank]count
type pairsHash map[int]int

// hasPairs checks if hash contains given number of pairs
func (p pairsHash) hasPairs(pair, count int) bool {
	pairsCount := 0
	for _, v := range p {
		if v == pair {
			pairsCount++
		}
	}
	return pairsCount == count
}

// newPairsHash counts pairs in hand
func newPairsHash(c []Card) pairsHash {
	pairs := make(pairsHash)
	for _, v := range c {
		if _, ok := pairs[v.Rank()]; ok {
			pairs[v.Rank()]++
		} else {
			pairs[v.Rank()] = 1
		}
	}
	return pairs
}

// Five cards of sequential rank and the same suit.
func StraightFlushEva(c []Card) Hand {
	if FlushEva(c) == Flush && StraightEva(c) == Straight {
		return StraightFlush
	}
	return NoHand
}

func FlushEva(c []Card) Hand {
	for i := 0; i < 4; i++ {
		if c[i].Suit() != c[i+1].Suit() {
			return NoHand
		}
	}
	return Flush
}

func FourOfAKindEva(c []Card) Hand {
	pairs := newPairsHash(c)
	if pairs.hasPairs(4, 1) && pairs.hasPairs(1, 1) {
		return FourOfAKind
	}
	return NoHand
}

func ThreeOfAKindEva(c []Card) Hand {
	pairs := newPairsHash(c)
	if pairs.hasPairs(3, 1) {
		return ThreeOfAKind
	}
	return NoHand
}

func FullHouseEva(c []Card) Hand {
	pairs := newPairsHash(c)
	if pairs.hasPairs(3, 1) && pairs.hasPairs(2, 1) {
		return FullHouse
	}
	return NoHand
}

func StraightEva(c []Card) Hand {
	sorted := make([]Card, len(c))
	copy(sorted, c)
	Sort(sorted)

	for i := 0; i < 4; i++ {
		if sorted[i+1].Rank() != sorted[i].Rank()-1 {
			return NoHand
		}
	}
	return Straight
}

func TwoPairEva(c []Card) Hand {
	pairs := newPairsHash(c)
	if pairs.hasPairs(2, 2) {
		return TwoPair
	}
	return NoHand
}

func OnePairEva(c []Card) Hand {
	pairs := newPairsHash(c)
	if pairs.hasPairs(2, 1) {
		return OnePair
	}
	return NoHand
}

func NoHandEva(c []Card) Hand {
	return NoHand
}

var evaluators = []Evaluator{
	StraightFlushEva,
	FourOfAKindEva,
	FullHouseEva,
	FlushEva,
	StraightEva,
	ThreeOfAKindEva,
	TwoPairEva,
	OnePairEva,
	NoHandEva,
}

// Evaluate takes array of N cards and returns best possible hand.
func Evaluate(cards []Card) Hand {
	var hands []Hand
	for _, f := range evaluators {
		hands = append(hands, f(cards))
	}

	var best Hand
	for _, h := range hands {
		if h > best {
			best = h
		}
	}

	return best
}
