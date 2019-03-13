package main

// . Khái niệm tương đương CLASS
// . co the compare 2 struct neu ben trong no la lieu du lieu co the so sanh (int, string, float...)
//   khong the compare 2 struct neu ben trong no la lieu du lieu khong the so sanh (map...)
// . zero value cua struct khong phai nil, thuoc tinh cua struct se la zero value cua chinh no

import (
	"fmt"
)

// 1a. tao struct
type Person struct {
	id   int
	name string
	age  int
}
type Animal struct {
	id   int
	name string
}

// 2. ke thua Struct
type Student struct {
	Person
	email string
}

// 4. struct long struct - nested struct
type Teacher struct {
	info  Person
	email string
	class int
}

// 1b. method struct
func (p Person) getId() (int, error) {
	return p.id, nil
}

// 1c. override
func (p Student) getId() (int, error) {
	p.Person.getId()

	fmt.Println("finish")
	return 0, nil
}

// 1. tao struct
//    method struct
//	  override
// 2. co ke thua
// 3. dinh nghia anonymous
// 4. struct long struct - nested struct: giong nhu has-a trong java
// 5. dinh nghia anonymous co ke thua
// 6. khai bao anonymous field
// 7. compare 2 struct
func main() {
	// 1a. tao struct
	per1 := Person{id: 1, name: "thien0", age: 24}
	// per1 := new(Person)

	// 1b. method struct
	fmt.Println(per1.getId())

	// 2.
	// st1 := Student{Person{1, "thien0", 24}, ""}
	st1 := Student{Person{id: 1, name: "thien1", age: 45}, "hieuthien95@gmail.com"}
	fmt.Println(st1)
	fmt.Println(st1.name)

	// 3. dinh nghia anonymous
	st2 := struct {
		id   int
		name string
	}{1, "thien2"}
	fmt.Println(st2)

	// 4. struct long struct - nested struct
	st3 := Teacher{
		Person{id: 1, name: "thien3", age: 45},
		"hieuthien95@gmail.com", 12,
	}

	fmt.Println(st3)
	fmt.Println(st3.info.name)

	// 5. dinh nghia anonymous co ke thua
	st4 := struct {
		Person
		email string
	}{
		Person{id: 1, name: "thien4", age: 24},
		"hieuthien95@gmail.com",
	}
	fmt.Println(st4)
	fmt.Println(st4.name)

	// 6. anonymous field
	dog1 := Animal{1, "dog"}
	fmt.Println(dog1)

	// 7. compare 2 struct co chung thuoc tinh so sanh duoc
	dog2 := Animal{1, "dog"}
	fmt.Println(dog1 == dog2)
}
