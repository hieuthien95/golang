package main

// . map là ref type
// . khong the so sanh == 2 map với nhau
//   chỉ có thể so sánh == nil
// . map nil khong the add key-value

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// 1. khai bao khong khoi tao
// 2. khai bao co khoi tao
// 3. khai bao bang map co gia tri mac dinh
// 4. thao tac voi map
//	  put
//	  get
// 	  delete
// 5. map là ref type
// 6. khong the so sanh == 2 map với nhau,
// 7. chỉ có thể so sánh == nil

func main() {
	// 1.
	var map1 map[string]int

	// 2.
	map2 := make(map[string]int)

	// 6.
	// 7.
	// fmt.Println(map1 == map2) error
	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)

	// 3.
	map3 := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	pp.Println(map3)

	// 4.
	m := make(map[string]int)

	// 4a. put
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// 4b. get
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// 4c. delete
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 5. map la ref type
	map4 := map[string]int{"key1": 1, "key2": 2, "key3": 3}
	map44 := map4
	map44["key2"] = 222
	pp.Println("ref type", map4)
}
