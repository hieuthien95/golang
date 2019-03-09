package main

// . Value Receiver: khong ref
// . Pointer Receiver: ref
// . non-struct

import "fmt"

type Shape struct {
	dientich int
	name     string
}

// 1. Value Receiver
func (sh Shape) changeNameByVal() {
	sh.name = "Not change name"
}

// 2. Pointer Receiver
func (sh *Shape) changeNameByPointer() {
	sh.name = "Changed name"
}

// 3. non-struct
type MyString string

func (s MyString) changeMyStringByVal() {
	s += "newwwww"
}
func (s *MyString) changeMyStringByPointer() {
	*s += "newwwww"
}

// 1. Value Receiver
// 2. Pointer Receiver
// 3. non-struct
func main() {
	sh := Shape{20, "Hinh CN"}
	fmt.Println(sh)

	// 1. Value Receiver
	sh.changeNameByVal()
	fmt.Println(sh)

	// 2. Pointer Receiver
	sh.changeNameByPointer()
	fmt.Println(sh)

	// 3. non-struct
	ms := MyString("my string")
	fmt.Println(ms)
	ms.changeMyStringByVal()
	fmt.Println(ms)
	ms.changeMyStringByPointer()
	fmt.Println(ms)
}
