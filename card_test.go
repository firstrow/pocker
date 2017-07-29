package pocker

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	var c Card
	c = NewCard("9C")
	if c.Rank() != 9 {
		t.Errorf("Expected 9. Got %d", c.Rank())
	}
	if c.Suit() != Clubs {
		t.Errorf("Expected clubs. Got %d", c.Suit())
	}

	c = NewCard("TD")
	if c.Rank() != 10 {
		t.Errorf("Expected 10. Got %s", c.Rank())
	}
	if c.Suit() != Diamonds {
		t.Errorf("Expected diamonds. Got %s", c.Suit())
	}
}

func TestCardsSort(t *testing.T) {
	cards := []Card{
		NewCard("5C"),
		NewCard("2D"),
		NewCard("JS"),
		NewCard("TH"),
		NewCard("KC"),
		NewCard("QS"),
	}
	Sort(cards)

	if cards[5].Rank() != 2 {
		t.Errorf("Expected 2. Got %d", cards[5].Rank())
	}
	if cards[4].Rank() != 5 {
		t.Errorf("Expected 5. Got %d", cards[4].Rank())
	}
	if cards[3].Rank() != 10 {
		t.Errorf("Expected 10. Got %d", cards[3].Rank())
	}
	if cards[2].Rank() != 11 {
		t.Errorf("Expected 11. Got %d", cards[2].Rank())
	}
	if cards[1].Rank() != 12 {
		t.Errorf("Expected 12. Got %d", cards[1].Rank())
	}
	if cards[0].Rank() != 13 {
		t.Errorf("Expected 13. Got %d", cards[0].Rank())
	}
}
