package main

func main() {
	deck := newDeck()

	// cards.saveToFile("my_cards")
	// deck := readFromFile("my_cards")

	deck.shuffle()
	deck.print()
}
