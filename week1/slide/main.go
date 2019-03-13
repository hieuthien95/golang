package main

// https://github.com/golang/go/wiki/SliceTricks

// . SLIDE là REF TYPE
// . length: so luong phan tu trong slide
//   capacity: số lượng phần tử còn lại trong array[], tính từ vị trí index của slide

import (
	"fmt"
)

// 1. khai bao basic
// 2. tạo slide: cắt tu 1 array
// 3. tao slide: cat tu 1 slide khac
// 4. slide la 1 ref type
// 5. length, capacity la gi?
// 6. MAKE
// 	  COPY: clone (value type)
//    APPEND
//    DELETE
func main() {
	// 1. khai bao basic
	fmt.Println("khai báo slide = [] không có số lượng")
	slide1 := []int{1, 2, 3}
	// var slide1 = new([50]int)[0:10]
	fmt.Println(slide1)

	// 2. tạo slide: cắt tu 1 array
	fmt.Println()
	fmt.Println("slide cắt từ array, slide khác:")
	array2 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	slide2 := array2[2:9] // [3 4 5 6 7 8 9]
	fmt.Println(slide2)
	slide2 = array2[2:6] // [3 4 5 6]
	fmt.Println(slide2)
	// [all]
	slide2 = array2[:]
	fmt.Println(slide2)
	//
	slide2 = array2[3:]
	fmt.Println(slide2)
	//
	slide2 = array2[:3]
	fmt.Println(slide2)

	// 3. tao slide: cat tu 1 slide khac
	fmt.Println()
	fmt.Println("slide in slide:")
	array5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slide5 := array5[2:5]
	slide55 := slide5[3:5]
	fmt.Println(array5)
	fmt.Println(slide5)
	fmt.Println(slide55)

	// 4. slide la 1 ref type
	fmt.Println()
	fmt.Println("demo ref:")
	array3 := [5]int{1, 2, 3, 4, 5}
	slide3 := array3[2:]
	slide3[0] = 999
	fmt.Println(array3)

	// 5. length, capacity la gi?
	fmt.Println()
	fmt.Println("len và cap:")
	array4 := [7]int{1, 2, 3, 4, 5, 6, 7}
	slide4 := array4[2:5]
	// slide4 = array4[:5] 	cap = 7
	// slide4 = array4[1:6] cap = 6
	// slide4 = array4[2:] 	cap = 5
	// slide4 = array4[6:7] cap = 1
	fmt.Printf("Array: %v \nSlide: %v \n  length: %v \n  capacity: %v",
		array4, slide4, len(slide4), cap(slide4))

	// 6. a. MAKE
	//    b. COPY: clone (value type)
	//    c. APPEND
	//    d. DELETE
	fmt.Println()
	fmt.Println("MAKE, COPY, APPEND:")
	// 6a.
	slide6 := make([]int, 2, 5) // [0 0]
	// 6c.
	slide6 = append(slide6, 100)       // [0 0 100]
	slide6 = append(slide6, 1, 2, 3)   // [0 0 100 1 2 3 4]
	slide6 = append(slide6, slide6...) // [0 0 100 1 2 3  0 0 100 1 2 3]
	fmt.Printf("value: %v, length: %v, capacity: %v", slide6, len(slide6), cap(slide6))
	fmt.Println()
	fmt.Println()
	// 6b.
	slide66 := make([]int, 3)
	copy(slide66, slide6)
	slide66[0] = 111
	fmt.Println(slide66)
	// 6d.
	src := []string{"thien", "quan", "nhan", "trang", "toan", "huy"}
	src = append(src[:1], src[2:]...)
	fmt.Println(src)

}
