package main

import "fmt"

func test() {
	defer func() {
		err := recover()
		fmt.Println("recover: stopped", err)
	}()
	
    panic(fmt.Sprintf("I am a  %d"))
}

func main() {
	test()
	fmt.Print("aaaa")
}
