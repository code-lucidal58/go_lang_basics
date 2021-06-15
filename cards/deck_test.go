package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 20
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(d))
	}
	firstElement := "Ace of Spades"
	if d[0] != firstElement {
		t.Errorf("Expected first card as %s, got %s", firstElement, d[0])
	}
	lastElement := "Five of Clubs"
	if d[len(d)-1] != lastElement {
		t.Errorf("Expected last card as %s, got %s", lastElement, d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	expectedLength := 20
	if len(loadedDeck) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(loadedDeck))
	}
	os.Remove("_decktesting")
}
