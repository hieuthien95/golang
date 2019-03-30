package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	y := v.Interface().(MyInt) // y will have type float64.
	fmt.Println(y)
}
