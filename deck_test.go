package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected: 16\nGot%v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected: 'Ace of Spades'\nGot: %v", d[0])
	}

	if (d[len(d)-1] != "Four of Clubs") {
		t.Errorf("Expected: 'Four of Clubs'\nGot: %v", d[0])
	}
}

func TestSaveToDeckAndReadDeckFile(t *testing.T) {
	testfile := "_decktesting"
	os.Remove(testfile)

	d := newDeck()
	d.saveToFile(testfile)

	loadedDeck, _ := readFile(testfile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected: 16\nGot%v", len(d))
	}

	os.Remove(testfile)
}