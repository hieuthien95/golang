package main

import (
	"fmt"
	"reflect"
)

// User ...
type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	School []string `json:"school"`
	ArrInt []int    `json:"arrInt"`
}

func main() {
	var user2 interface{}
	user2 = User{Name: "Thien", Age: 18, School: []string{"XD", "NDC", "TDT"}, ArrInt: []int{1, 2, 3}}
	fmt.Println(user2)

	fmt.Println("======================================================")
	user1 := []interface{}{
		"Thien",
		18,
		[]string{"XD", "NDC", "TDT"},
		[]int{1, 2, 3}, func() {},
		User{"Thien", 18, []string{"XD", "NDC", "TDT"}, []int{1, 2, 3}},
		user2,
	}
	fmt.Println(user1)

	// 1. Append | ValueOf | Indirect
	a := reflect.Indirect(reflect.ValueOf(user1[5]))
	fmt.Println(a)
	usr1 := reflect.Indirect(reflect.ValueOf(user1))
	usr1 = reflect.Append(usr1, a)

	fmt.Println(usr1)

	// 2. Kind
	for _, v := range user1 {
		switch vl := reflect.ValueOf(v); vl.Kind() {
		case reflect.String:
			fmt.Println("+string")
		default:
			fmt.Println(vl.Kind())
		}
	}
}
