package util

import (
	"flag"
	"fmt"
)

func Flag() {
	str1 := flag.String("u", "root", "username")
	str2 := flag.String("p", "", "password")
	flag.Parse()
	fmt.Println("Username : ", *str1)
	fmt.Println("Password : ", *str2)
}
