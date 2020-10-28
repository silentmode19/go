/* Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	a := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}
	b := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(a.first)
	fmt.Println(a.last)
	for i, v := range a.favFlavors {
		fmt.Println(i, v)
	}
	fmt.Println(b.first)
	fmt.Println(b.last)
	for i, v := range b.favFlavors {
		fmt.Println(i, v)
	}

}
