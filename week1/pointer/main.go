package main

// pointer la ref type

import "fmt"

// ref type
func funPointer1(poi *int) {
	*poi = 111
}

// value type
func funNum1(num int) {
	num = 222
}

// 1. poiner normal
// 2. pointer arrays
func main() {

	// 1. poiner normal
	var number int = 12

	// var poi *int 	// zero value = nil
	var poi = new(int) // zero &value = hexa, *value = 0
	fmt.Println("zero: ", &poi, ": ", *poi)
	poi = &number
	fmt.Println(*poi)

	// ref type
	funPointer1(poi)
	fmt.Println("ref type", *poi)

	// value type
	funNum1(number)
	fmt.Println("value type", number)

	// 2. pointer arrays
	array := [4]int{1, 2, 3, 4}
	// poi2 := &array
	var poi2 *[4]int
	// var poi2 *[3]int err
	// var poi2 *[5]int err
	poi2 = &array
	fmt.Printf("address: %v \nvalue: %v", poi2, *poi2)
}
