package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doSth1(i int) {
	fmt.Println(i)
	for i := 0; i < 100000; i++ {
		for j := 0; j < 100; j++ {
		}
	}
}

func doSth2(i int, groupTest *sync.WaitGroup) {
	fmt.Println(i)
	for i := 0; i < 100000; i++ {
		for j := 0; j < 100; j++ {
		}
	}
	groupTest.Done()
}

func main() {
	fmt.Println("BEGIN")

	timenow := time.Now()
	defer func(start time.Time) {
		ellapsed := time.Since(timenow)
		fmt.Println(ellapsed)
	}(timenow)

	// test none gorountine
	// for i := 1; i <= 7; i++ {
	// 	doSth1(i)
	// }

	// // test gorountine
	// groupTest := new(sync.WaitGroup)
	// groupTest.Add(7)
	// for i := 1; i <= 7; i++ {
	// 	go doSth2(i, groupTest)
	// }
	// groupTest.Wait()

	// test parellellism
	groupTest := new(sync.WaitGroup)
	runtime.GOMAXPROCS(4)
	groupTest.Add(7)
	for i := 1; i <= 7; i++ {
		go doSth2(i, groupTest)
	}
	groupTest.Wait()

	fmt.Println("END")
}
