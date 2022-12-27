package main

import (
	"fmt"
	"net/http"
	"time"
)

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

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()

	// cardss := newDeckFromFile("my_cards")
	// cardss.shuffle()
	// cardss.print()
	// println(cards.toString())

	links := []string{
		"https://www.google.com.vn/",
		"https://www.youtube.com/",
		"https://www.facebook.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	second := time.Now()
	fmt.Println((second))
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
	// fmt.Printf("%+v", p)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
