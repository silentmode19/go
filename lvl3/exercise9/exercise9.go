//Create a program that uses a switch statement with the switch expression
//specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import (
	"fmt"
)

func main() {
	favSport := "Swimming"
	switch favSport {
	case "Soccer":
		fmt.Println("Where's the ball?")
	case "Tennis":
		fmt.Println("Where's the court?")
	case "Swimming":
		fmt.Println("Let's go!")
	}
}
