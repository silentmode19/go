/* create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
*/

package main

import (
	"fmt"
)

func main() {
	x := foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	y := bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x, y)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
