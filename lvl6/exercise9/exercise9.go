/* A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument */

package main

import "fmt"

func main() {
	factorial(4)
	fmt.Println(4)
}
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
