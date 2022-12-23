package main

func main() {

	cardss := newDeckFromFile("my_cards")
	cardss.shuffle()
	cardss.print()
	// println(cards.toString())
}
