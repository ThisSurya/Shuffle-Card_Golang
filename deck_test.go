package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Error expected found 16, in real: %v", len(d))
	}

	if d[0] != "satu dari Heart" {
		t.Errorf("Error expected found satu dari Heart, in real: %v ", d[0])
	}

	if d[len(d)-1] != "empat dari Club" {
		t.Errorf("Error expected found empat dari Club, in real: %v ", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	cards := newDeckfromFile("_decktesting")

	if len(cards) != 16 {
		t.Errorf("Error expected deck card 16, found %v", len(cards))
	}
	os.Remove("_decktesting")
}
