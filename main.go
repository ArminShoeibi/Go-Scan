package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string

	fmt.Print("Please enter you first name: ")
	_, _ = fmt.Scan(&firstName)

	fmt.Print("Please enter you last name: ")
	_, _ = fmt.Scan(&lastName)

	fmt.Println("Hello, Welcome", firstName, lastName)

}
