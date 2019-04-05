package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("expected deck length of 52, but got %v", len(d))
	}

	first, last := d[0], d[len(d)-1]

	if first != "Ace of Spades" {
		t.Errorf("expected card is `Ace of Spades`, but got %v", first)
	}
	if last != "King of Clubs" {
		t.Errorf("expected card is `Ace of Spades`, but got %v", last)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()

	d.saveToFile("_decktesting")
	readDeck := readFromFile("_decktesting")

	if len(d) != len(readDeck) {
		t.Errorf("Expected %v cards in deck, got %v", len(d), len(readDeck))
	}

	os.Remove("_decktesting")
}
