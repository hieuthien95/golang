package main

// . bool
// . string
// . int  int8  int16  int32  int64
//   uint uint8 uint16 uint32 uint64 uintptr
// . byte // alias for uint8
// . rune // alias for int32
//        // represents a Unicode code point
// . float32 float64

import "fmt"

// . complex64 complex128

// 1. khai bao basic
// 2. Convert
// 	  chi co the convert tu kieu du lieu lon => nho
// 3. RUNE
// 4. INT32
func main() {
	// 1. khai bao basic
	var a int32
	var b int
	var c float32
	var d complex64
	var e string

	a = 1
	b = 2
	c = 2.1
	e = "nordic"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// 2. Convert
	var a2 = 1
	f2 := int(a2)
	fmt.Println(f2)

	// 1. RUNE
	fmt.Println("RUNE: ")
	myString := "Thiện"
	runes := []rune(myString)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	// 2. INT32
	fmt.Println()
	fmt.Println("INT32: ")
	myStrings := "Thiện"
	runess := []int32(myStrings)
	for i := 0; i < len(runess); i++ {
		fmt.Printf("%c", runess[i])
	}
}
