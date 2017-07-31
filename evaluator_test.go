package pocker

import (
	"testing"
)

func TestStraightFlushSuccess(t *testing.T) {
	cards := CodesToCards([]string{"QC", "KC", "JC", "TC", "9C"})
	r := StraightFlushEva(cards)

	if r != StraightFlush {
		t.Errorf("StraightFlush was expected. Got %s", r)
	}
}

func TestStraightFlushSuccess2(t *testing.T) {
	cards := CodesToCards([]string{"JS", "9S", "TS", "8S", "7S"})
	r := StraightFlushEva(cards)

	if r != StraightFlush {
		t.Errorf("StraightFlush was expected. Got %s", r)
	}
}

func TestStraightFlushFail(t *testing.T) {
	cards := CodesToCards([]string{"JS", "TS", "9S", "6S", "7C"})

	r := StraightFlushEva(cards)

	if r != NoHand {
		t.Errorf("NoHand was expected. Got %s", r)
	}
}

func TestStraightSuccess(t *testing.T) {
	cards := CodesToCards([]string{"JS", "TD", "8S", "7C", "9S"})

	r := StraightEva(cards)

	if r != Straight {
		t.Errorf("Straight was expected. Got %s", r)
	}
}

func TestStraightFail(t *testing.T) {
	cards := CodesToCards([]string{"JS", "TD", "QS", "8S", "7C"})

	r := StraightEva(cards)

	if r != NoHand {
		t.Errorf("NoHand was expected. Got %s", r)
	}
}

func TestFullHouse(t *testing.T) {
	cards := CodesToCards([]string{"KC", "KD", "KH", "QS", "QC"})

	r := FullHouseEva(cards)

	if r != FullHouse {
		t.Errorf("FullHouse was expected. Got %s", r)
	}
}

func TestFourOfAKind(t *testing.T) {
	cards := CodesToCards([]string{"5C", "5D", "5H", "5S", "QC"})

	r := FourOfAKindEva(cards)

	if r != FourOfAKind {
		t.Errorf("FourOfAKind was expected. Got %s", r)
	}
}

func TestThreeOfAKind(t *testing.T) {
	cards := CodesToCards([]string{"5C", "5D", "5H", "KS", "QC"})

	r := ThreeOfAKindEva(cards)

	if r != ThreeOfAKind {
		t.Errorf("ThreeOfAKind was expected. Got %s", r)
	}
}

func TestTwoPair(t *testing.T) {
	cards := CodesToCards([]string{"5C", "5D", "6H", "6S", "QC"})

	r := TwoPairEva(cards)

	if r != TwoPair {
		t.Errorf("TwoPair was expected. Got %s", r)
	}
}

func TestOnePair(t *testing.T) {
	cards := CodesToCards([]string{"5C", "5D", "3H", "4S", "QC"})

	r := OnePairEva(cards)

	if r != OnePair {
		t.Errorf("OnePair was expected. Got %s", r)
	}
}

func TestPairsHash(t *testing.T) {
	cards := CodesToCards([]string{"QS", "QD", "QC", "8S", "8C"})

	pairs := newPairsHash(cards)
	if pairs.hasPairs(2, 1) == false {
		t.Error("Pair of two was expected")
	}
	if pairs.hasPairs(3, 1) == false {
		t.Error("Pair of three was expected")
	}
}

func TestPairsHashFail(t *testing.T) {
	cards := CodesToCards([]string{"QS", "KD", "AC", "2S", "5C"})

	pairs := newPairsHash(cards)
	if pairs.hasPairs(2, 1) == true {
		t.Error("No pairs was expected")
	}
}

func TestEvaluateStraightFlush(t *testing.T) {
	cards := CodesToCards([]string{"KC", "QC", "JC", "TC", "9C"})

	hand := Evaluate(cards)
	if hand != StraightFlush {
		t.Error("StraightFlush was expected")
	}
}

func TestEvaluateFourOfAKind(t *testing.T) {
	cards := CodesToCards([]string{"KC", "KD", "KS", "KH", "9C"})

	hand := Evaluate(cards)
	if hand != FourOfAKind {
		t.Error("StraightFlush was expected")
	}
}

func TestEvaluateFullHouse(t *testing.T) {
	cards := CodesToCards([]string{"3C", "3D", "3S", "9H", "9C"})

	hand := Evaluate(cards)
	if hand != FullHouse {
		t.Error("FullHouse was expected")
	}
}

func TestBestHandStraightFlush(t *testing.T) {
	hand := []string{"TH", "JH", "QC", "QD", "QS"}
	deck := []string{"QH", "KH", "AH", "2S", "6S"}

	r := BestHand(hand, deck)
	if r != StraightFlush {
		t.Errorf("StraightFlush expected. Got %s", r)
	}
}

func TestBestHandStraight(t *testing.T) {
	hand := []string{"AC", "2D", "6C", "3S", "KD"}
	deck := []string{"5S", "4D", "KS", "AS", "4C"}

	r := BestHand(hand, deck)
	if r != Straight {
		t.Errorf("Straight expected. Got %s", r)
	}
}
