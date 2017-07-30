package pocker

import ()

type Hand int
type Evaluator func([]Card) Hand

const (
	NoHand = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	Straight // Done
	Flush    // Done
	FullHouse
	FourOfAKind
	StraightFlush //Done
)

// Five cards of sequential rank and the same suit.
func StraightFlushEva(c []Card) Hand {
	if FlushEva(c) != Flush {
		return NoHand
	}

	if StraightEva(c) != Straight {
		return NoHand
	}

	return StraightFlush
}

func FlushEva(c []Card) Hand {
	for i := 0; i < 4; i++ {
		if c[i].Suit() != c[i+1].Suit() {
			return NoHand
		}
	}
	return Flush
}

func StraightEva(c []Card) Hand {
	for i := 0; i < 4; i++ {
		if c[i+1].Rank() != c[i].Rank()-1 {
			return NoHand
		}
	}
	return Straight
}

func HighCardEva(cards []Card) Hand {
	return HighCard
}

func OnePairEva(cards []Card) Hand {
	return OnePair
}

var evaluators = []Evaluator{
	StraightFlushEva,
	HighCardEva,
	OnePairEva,
}

// Evaluate takes array of 5 cards and returns best possible hand.
func Evaluate(cards []Card) Hand {
	var r = make([]Hand, 5)
	for _, f := range evaluators {
		r = append(r, f(cards))
	}
	// TODO: sort r and return highest
	return 1
}
