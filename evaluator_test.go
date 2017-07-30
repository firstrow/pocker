package pocker

import (
	"testing"
)

func TestStraightFlushSuccess(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("KC"),
		NewCard("QC"),
		NewCard("JC"),
		NewCard("TC"),
		NewCard("9C"),
	}
	Sort(cards)

	r := StraightFlushEva(cards)

	if r != StraightFlush {
		t.Errorf("StraightFlush was expected. Got %s", r)
	}
}

func TestStraightFlushSuccess2(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("JS"),
		NewCard("TS"),
		NewCard("9S"),
		NewCard("8S"),
		NewCard("7S"),
	}
	Sort(cards)

	r := StraightFlushEva(cards)

	if r != StraightFlush {
		t.Errorf("StraightFlush was expected. Got %s", r)
	}
}

func TestStraightFlushFail(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("JS"),
		NewCard("TS"),
		NewCard("9S"),
		NewCard("6S"),
		NewCard("7C"),
	}
	Sort(cards)

	r := StraightFlushEva(cards)

	if r != NoHand {
		t.Errorf("NoHand was expected. Got %s", r)
	}
}

func TestStraightSuccess(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("JS"),
		NewCard("TD"),
		NewCard("9S"),
		NewCard("8S"),
		NewCard("7C"),
	}
	Sort(cards)

	r := StraightEva(cards)

	if r != Straight {
		t.Errorf("Straight was expected. Got %s", r)
	}
}

func TestStraightFail(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("JS"),
		NewCard("TD"),
		NewCard("QS"),
		NewCard("8S"),
		NewCard("7C"),
	}
	Sort(cards)

	r := StraightEva(cards)

	if r != NoHand {
		t.Errorf("NoHand was expected. Got %s", r)
	}
}

func TestFullHouse(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("KC"),
		NewCard("KD"),
		NewCard("KH"),
		NewCard("QS"),
		NewCard("QC"),
	}
	Sort(cards)

	r := FullHouseEva(cards)

	if r != FullHouse {
		t.Errorf("FullHouse was expected. Got %s", r)
	}
}

func TestFourOfAKind(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("5C"),
		NewCard("5D"),
		NewCard("5H"),
		NewCard("5S"),
		NewCard("QC"),
	}
	Sort(cards)

	r := FourOfAKindEva(cards)

	if r != FourOfAKind {
		t.Errorf("FourOfAKind was expected. Got %s", r)
	}
}

func TestThreeOfAKind(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("5C"),
		NewCard("5D"),
		NewCard("5H"),
		NewCard("KS"),
		NewCard("QC"),
	}
	Sort(cards)

	r := ThreeOfAKindEva(cards)

	if r != ThreeOfAKind {
		t.Errorf("ThreeOfAKind was expected. Got %s", r)
	}
}

func TestTwoPair(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("5C"),
		NewCard("5D"),
		NewCard("6H"),
		NewCard("6S"),
		NewCard("QC"),
	}
	Sort(cards)

	r := TwoPairEva(cards)

	if r != TwoPair {
		t.Errorf("TwoPair was expected. Got %s", r)
	}
}

func TestOnePair(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("5C"),
		NewCard("5D"),
		NewCard("3H"),
		NewCard("4S"),
		NewCard("QC"),
	}
	Sort(cards)

	r := OnePairEva(cards)

	if r != OnePair {
		t.Errorf("OnePair was expected. Got %s", r)
	}
}

func TestPairsHash(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("QS"),
		NewCard("QD"),
		NewCard("QC"),
		NewCard("8S"),
		NewCard("8C"),
	}
	Sort(cards)

	pairs := newPairsHash(cards)
	if pairs.hasPairs(2, 1) == false {
		t.Error("Pair of two was expected")
	}
	if pairs.hasPairs(3, 1) == false {
		t.Error("Pair of three was expected")
	}
}

func TestPairsHashFail(t *testing.T) {
	// TODO: Move cards creation and sort to reusable metod
	cards := []Card{
		NewCard("QS"),
		NewCard("KD"),
		NewCard("AC"),
		NewCard("2S"),
		NewCard("5C"),
	}
	Sort(cards)

	pairs := newPairsHash(cards)
	if pairs.hasPairs(2, 1) == true {
		t.Error("No pairs was expected")
	}
}
