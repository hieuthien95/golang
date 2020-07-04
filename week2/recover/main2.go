package main

import "fmt"

func test() (data string) {
    defer func() {
        err := recover()
        fmt.Println("recover: stopped", err)
        data = "catch return dc ne"
    }()
	
    panic(fmt.Sprintf("I am a  %d"))

    return "bbbb"
}

func main() {
	data:=test()
	fmt.Println("aaaa")
	fmt.Println("output: ", data)
}
