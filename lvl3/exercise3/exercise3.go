/* Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

package main

import (
	"fmt"
)

func main() {
	for i := 1991; i < 2021; i++ {
		fmt.Println(i)
	}
}
