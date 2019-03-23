package main

import "fmt"

func main() {
	var action int
	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println("I am a  Student")
	case 2:
		fmt.Println("I am a  Professional")
	default:
		panic(fmt.Sprintf("I am a  %d", action))
	}
	fmt.Println("")
	fmt.Println("Finish")

	// fmt.Println("Enter 1 for US and 2 for UK")
	// fmt.Scanln(&action)
	// /*  Use of Switch Case in Golang */
	// switch {
	// case 1:
	// 	fmt.Println("US")
	// case 2:
	// 	fmt.Println("UK")
	// default:
	// 	panic(fmt.Sprintf("I am a  %d", action))
	// }
}
