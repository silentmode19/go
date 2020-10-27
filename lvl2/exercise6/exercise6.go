//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import (
	"fmt"
)

const (
	year int = 2020
	//	_    int = year + iota
	a int = year + iota
	b int = year + iota
	c int = year + iota
	d int = year + iota
)

func main() {
	fmt.Println(a, b, c, d)
}
