package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Wrong deck length")
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

}
