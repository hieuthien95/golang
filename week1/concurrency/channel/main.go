package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go showInputChannel() khong waiting inputChannel() function finish,
// go showInputChannel() chi waiting den khi trong channel c co gia tri
func inputChannel(c chan string) {
	fmt.Println("<<inputChannel run>>")

	fmt.Print("Please type input Channel: ")
	var input string
	fmt.Scanln(&input)

	c <- input

	for i := 0; i <= 1000; i++ {
		if i%10 == 0 {
			fmt.Print(i, " ")
		}
		if i == 999 {
			fmt.Println("finish inputChannel", i)
		}
	}
}

// inputChannel() va showInputChannel() thuc hien cung luc
// nhung den xu ly "intput := <-c" se pause den khi nao trong channel c co gia tri
func showInputChannel(c chan string, groupTest *sync.WaitGroup) {
	fmt.Println("<<showInputChannel run>>")

	// pause den khi channel c co gia tri
	intput := <-c
	fmt.Println("\nShow Input Channel: ", intput)

	// cau lenh phia duoi Done() co the se khong chay, vi main() finish som hon
	// 			phia tren chay binh thuong
	groupTest.Done()

	// khong xu ly het
	fmt.Println("Finish showInputChannel")
	for i := 0; i <= 1000; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	fmt.Println("BEGIN")

	groupTest := new(sync.WaitGroup)
	runtime.GOMAXPROCS(1)
	groupTest.Add(1)

	var c = make(chan string)

	go inputChannel(c)
	go showInputChannel(c, groupTest)

	groupTest.Wait()

	fmt.Println("END")

}
