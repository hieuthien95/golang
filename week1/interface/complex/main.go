package main

// . liên quan tới Receiver
// . phải khai báo đủ method trong interface, thì mới xem như đã impl

// . Multiple interface: co the impl nhiu interface
// . Embed interface: kết hợp nhìu interface vao 1 interface mới
// . Empty interface: là generic type <T>

import (
	"fmt"
	"reflect"
)

// B1: interface
// 1.
type IStudent interface {
	learn()
}

// 2.
type ITeacher interface {
	teach()
}

// 3.
type IDr interface {
	IStudent
	ITeacher
	research()
}

// B2: tao struct
type Person struct {
	name string
}

// B3: tao "đầy đủ" các method
// 1. impl từ IAnimal
func (c Person) learn() {
	fmt.Println("learn...............")
}

// 2. impl them Dongvatable
func (c Person) teach() {
	fmt.Println("teach...............")
}

// 3. Embed 3 cái trên + method này
func (c Person) research() {
	fmt.Println("research...............")
}

// 4.
func getout(itf interface{}) {
	var typ_ Person
	boo := reflect.TypeOf(itf) == reflect.TypeOf(typ_)
	if boo == true {
		mydg := itf.(Person)
		fmt.Printf("He is [%T]: %v", mydg, mydg.name)
		fmt.Println()

		mydg.learn()
		mydg.teach()
		mydg.research()

	} else {
		fmt.Printf("Just is [%T]: %v", itf, itf)
		fmt.Println()
	}
}

// 1. impl 1 interface 			--> learn()
// 2. impl từ nhieu interface	--> có thêm teach()
// 3. Embed interface			--> learn() teach()
// 4. Empty interface
// 5. Pointer interface
func main() {

	// // 1.
	// var an1 IStudent
	// an1 = Person{"Poppy"}
	// an1.learn()

	// // 2.
	// fmt.Println()
	// var dv2 ITeacher
	// dv2 = Person{"Poppy"}
	// dv2.teach()

	// // 3.
	// fmt.Println()
	// var to3 IDr
	// to3 = Person{"Poppy"}
	// to3.learn()
	// to3.teach()
	// to3.research()

	// 4.
	var istudent IStudent = Person{"Mr Stud"}
	var iteach ITeacher = Person{"Mr Tee"}
	var idr IDr = Person{"Dr Tra"}
	fmt.Println()

	getout(10)
	getout("abc")

	getout(istudent)
	getout(iteach)
	getout(idr)
}
