//Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

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

	/* 	fmt.Println(a.first)
	   	fmt.Println(a.last)
	   	for i, v := range a.favFlavors {
	   		fmt.Println(i, v)
	   	}
	   	fmt.Println(b.first)
	   	fmt.Println(b.last)
	   	for i, v := range b.favFlavors {
	   		fmt.Println(i, v)
	   	} */

	m := map[string]person{
		a.last: a,
		b.last: b,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, v2 := range v.favFlavors {
			fmt.Println(i, v2)
		}
	}
}
