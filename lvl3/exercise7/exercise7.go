//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import (
	"fmt"
)

func main() {
	x := "what"
	if x == "true" {
		fmt.Println("True")
	} else if x == "maybe" {
		fmt.Println("Maybe")
	} else {
		fmt.Println("I dont know")
	}
}
