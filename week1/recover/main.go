package main

import "fmt"

func main() {
	defer func() {
		action := recover()
		fmt.Println("recover: stopped", action)
	}()

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
}
