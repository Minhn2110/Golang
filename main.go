package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	// cardss := newDeckFromFile("my_cards")
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// cardss.shuffle()
	// cardss.print()
	// println(cards.toString())
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
	// fmt.Printf("%+v", p)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
