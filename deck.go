package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { 
	return strings.Join([]string(d), ", ")
}

func (d deck) toByte() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
	err := os.WriteFile(filename, d.toByte(), os.ModeDevice)
	return err
}

func readFile(filename string) (deck, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	string := strings.Split(string(bytes), ", ")
	return deck(string), nil
}

func toDeck(byteSlice []byte) deck {
	stringDeck := strings.Split(string(byteSlice), ", ")
	return deck(stringDeck)
}

func (d deck) shuffle() {
	for i := range d {
		random := rand.Intn(len(d) - 1)
		d[i], d[random] = d[random], d[i]
	}
}