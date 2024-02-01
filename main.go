package main

import (
	"fmt"
)

func main() {

	var stop bool
	var d deck
	var filename string
	for stop != true {
		var choice string

		fmt.Println("Choose an option in the menu: \n[1] Create a new deck \n[2] Shuffle \n[3] Deal \n[4] Save deck to file \n[5] Retrieve deck from file\n[0] Exit")
		_, err := fmt.Scan(&choice)
		
		if err != nil {
			fmt.Println("Error: ", err)
		}

		switch choice {
		case "1":
			d = newDeck()
			d.print()
		case "2":
			if (d == nil) {
				optionBeforeDeckCreated()
				break
			}
			d.shuffle()
			d.print()
		case "3":
			if (d == nil) {
				optionBeforeDeckCreated()
				break
			}
			var handsize int
			fmt.Print("What is the hand size? ")
			_, err := fmt.Scan(&handsize)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			deal, remainingCards := deal(d, handsize)
			fmt.Println("Deal: ", deal)
			fmt.Println("Remaining cards: ", remainingCards)
		case "4":
			if (d == nil) {
				optionBeforeDeckCreated()
				break
			}
			fmt.Print("What is the file name? ")
			_, err := fmt.Scan(&filename)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			d.saveToFile(filename)
		case "5":
			if (filename == "") {
				optionBeforeFileSave()
				break
			}
			deckFromFile, err := readFile(filename)
			if (err != nil) {
				fmt.Println("Error: ", err)
				break
			}
			deckFromFile.print()
		case "0":
			stop = true
		default:
			invalidOption()
		}
	}
}

func optionBeforeDeckCreated() {
	fmt.Println("You need to create a deck before using this option.")
}

func optionBeforeFileSave() {
	fmt.Println("You need to save the file before retrieving it.")
}

func invalidOption() {
	fmt.Println("You need to input a valid option.")
}