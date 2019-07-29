package main

func main() {
	cards := newDeck()
	cards.saveToFile("My_Cards")
	cards.print()
}
