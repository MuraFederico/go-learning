package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	length := len(deck)

	if length != 52 {
		t.Errorf("Expected deck length of 52, but got %v", length)
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%v'", deck[0])
	}

	if deck[length-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got '%v'", deck[length-1])
	}
}
