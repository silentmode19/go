//get this code working:https://play.golang.org/p/j-EA6003P0

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
