/* Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/

package main

import (
	"fmt"
)

func main() {
	y := 1991
	for {
		if y > 2020 {
			break
		}
		fmt.Println(y)
		y++
	}
}
