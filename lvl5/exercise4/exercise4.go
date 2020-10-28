//Create and use an anonymous struct.

package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first string
		last  string
		hobby string
	}{
		first: "Alfred",
		last:  "Pennyworth",
		hobby: "Cleaning up",
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.hobby)
}
