/* Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.
*/

package main

import (
	"fmt"
)

func main() {
	a := (123 == 123)
	b := (123 <= 124)
	c := (123 >= 122)
	d := (123 != 125)
	e := (123 < 126)
	f := (123 > 121)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
