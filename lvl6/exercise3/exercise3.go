//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("first?or deferred?")
}
func bar() {
	fmt.Println("bar")
}
