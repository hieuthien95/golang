package main

// . ARRAY la value type
//   khong phai ref type

import "fmt"

// 1. khai bao array basic
// 2. khai bao khong biet so luong phan tu
// 3. duyet for array thao cach RANGE
func main() {
	// 1. khai bao array basic
	var b [5]int
	b[1] = 100
	fmt.Println(b)

	var myArr1 = [5]int{100}
	fmt.Printf("%v %v", len(myArr1), myArr1)

	// 2. khai bao khong biet so luong phan tu
	fmt.Println()
	myArr2 := [...]int{1, 2, 3, 4, 5}
	myArr2[4] = 555
	// myArr2[6] = 555 error
	fmt.Printf("%v %v", len(myArr2), myArr2)

	// 3. duyet for array thao cach RANGE
	fmt.Println()
	fmt.Println("for array, dung cach 2")
	myArr3 := [...]string{"thien", "quan", "nhan", "trang", "toan", "huy"}
	for _, value := range myArr3 {
		fmt.Print(value, " ")
	}
	fmt.Println()
	for index, value := range myArr3 {
		fmt.Print(index+1, ".", value, " ")
	}
}
