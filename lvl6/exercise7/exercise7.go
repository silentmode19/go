//Assign a func to a variable, then call that func

package main

import (
	"fmt"
)

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)
}
