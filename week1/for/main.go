package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	// a := []int{1, 2, 3, 4, 5, 6}

	// for
	// for i := 0; i < 5; i++ {
	// 	pp.Println(a[i])
	// }

	// while
	// for {
	// 	pp.Println(1)
	// }

	// for range
	// for i := range a {
	// 	pp.Print(i)
	// }
	// for i, v := range a {
	// 	pp.Print(i, " - ")
	// 	pp.Println(v)
	// }
	// for i, _ := range a {
	// 	pp.Print(i)
	// }
	// for _, v := range a {
	// 	pp.Print(v)
	// }

	m := map[string]int{"key1": 1, "key2": 2}
	// _,v k,_
	pp.Println(m)
	for k, v := range m {
		pp.Print(k, "-")
		pp.Println(v)
	}
}
