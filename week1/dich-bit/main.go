package main

import "fmt"

func main() {
	var x, y, z int
	x = 75
	y = 25
	z = x & y // &    bitwise AND            integers
	fmt.Println(z)
	z = x | y // |    bitwise OR             integers
	fmt.Println(z)
	z = x ^ y // ^    bitwise XOR            integers
	fmt.Println(z)
	z = x &^ y // &^   bit clear (AND NOT)    integers
	fmt.Println(z)
}
