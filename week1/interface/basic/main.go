package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type IPrinter interface {
	print()
	printInfo()
}

type Person struct {
	id   int
	name string
	age  int
}

type Animal struct {
	id   int
	name string
}

func (p Person) print() {
	fmt.Println("print:", p.id, "-", p.name)
}

func (p Animal) print() {
	fmt.Println("print:", p.id, "-", p.name)
}

func (p Person) printInfo() {
	fmt.Println("printInfo:", p.id, "-", p.name)
}

func doSth(i interface{}) {
	pp.Println(i)
}

func main() {

	var a, p IPrinter
	p = Person{1, "thien", 24}
	// a = Animal{1, "animal"} // Animal does not implement IPrinter (missing printInfo method)

	p.print()
	p.printInfo()

	doSth(p)
	doSth(a)
}
