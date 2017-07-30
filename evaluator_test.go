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
