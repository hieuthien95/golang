package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// var x bool = true
	// var y bool
	// var z bool

	// if x {
	// 	pp.Println(true)
	// } else {
	// 	pp.Println(false)
	// }

	// if x || (y && z) {
	// 	pp.Println(true)
	// }

	// =============================
	// switch x {
	// case true:
	// 	pp.Println(true)
	// case false:
	// 	pp.Println(false)
	// default:
	// 	pp.Println("true/false")
	// }

	a := 1
	switch {
	case a == 1:
		pp.Println(1)
	case a == 2:
		pp.Println(2)
	}

}
