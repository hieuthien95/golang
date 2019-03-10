package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// 1. function normal
func add(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, errors.New("mat vui roi")
	}
	return a + b, nil
}
func sayHi(s string) {
	fmt.Println("Hi", s, "!")
}

// 2a. Variadic Function: neu truyen vao array param thi la value type
// 2b. 					  neu truyen vao slide..., thì nó là ref type
// 2c. 					  different types of arguments
func demoVariadicFunction(item string, list ...int) {
	list[0] = 999
	fmt.Println(item, list)
}
func variadicDiffType(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

// 3. function Values
func funValues(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 4. function Closures
func adderFunClosures() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 1. function normal
// 2. Variadic Function: neu truyen vao array param thi la value type
//   					 neu truyen vao slide..., thì nó là ref type
// 3. function Values
// 4. function Closures
func main() {
	// 1. function normal
	sayHi("Thiện")
	//
	sum, err := add(0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}

	// 2a. neu truyen vao array param thi la value type
	demoVariadicFunction("echo:", 1, 2, 3, 4, 5, 6, 7)
	array := []int{1, 2, 3}
	// 2b. neu truyen vao slide..., thì nó là ref type
	demoVariadicFunction("echo:", array...)
	fmt.Println(array)
	// 2c.
	variadicDiffType(1, "red", true, 1.5, []string{"foo", "bar"}, map[string]int{"apple": 1, "tomato": 2})

	// 3. function Values
	tenFuncVal := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(tenFuncVal(5, 12))

	fmt.Println(funValues(tenFuncVal))
	fmt.Println(funValues(math.Pow))

	// 4. function Closures
	// funcName := adderFunClosures()

	// num := funcName(1)
	// num = funcName(2)
	// num = funcName(3)

	// fmt.Println(num)
}
