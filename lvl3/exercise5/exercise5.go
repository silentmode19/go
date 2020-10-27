//Print out the remainder (modulus)
//which is found for each number between 10 and 100 when it is divided by 4.

package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 101; i++ {
		fmt.Printf("Number: %v, Modulus: %v\n", i, i%4)
	}
}
