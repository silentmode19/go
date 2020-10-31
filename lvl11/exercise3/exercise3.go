//Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.

package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	log.Println("foo ran ", e)
}
