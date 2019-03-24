package main

import (
	"fmt"
	"reflect"
)

func Contains(valInput interface{}, arrInput ...interface{}) bool {
	for _, value := range arrInput {
		// C1:
		// switch {
		// case reflect.TypeOf(value) == reflect.TypeOf([]int64{}):
		// 	for _, val := range value.([]int) {
		// 		if val == valInput {
		// 			return true
		// 		}
		// 	}
		// case reflect.TypeOf(value) == reflect.TypeOf([]string{}):
		// 	for _, val := range value.([]string) {
		// 		if val == valInput {
		// 			return true
		// 		}
		// 	}
		// case reflect.TypeOf(value) == reflect.TypeOf([]float32{}):
		// 	// TODO
		// case reflect.TypeOf(value) == reflect.TypeOf([]int{}):
		// 	// TODO
		// }

		// C2:
		switch value.(type) {
		case []int:
			for _, val := range value.([]int) {
				if val == valInput {
					return true
				}
			}
		case []int32:
			// TODO
		case []int64:
			// TODO
		case []float32:
			// TODO
		case []float64:
			// TODO

		case [5]string:
			for _, val := range value.([5]string) {
				if val == valInput {
					return true
				}
			}
		}

		// C3:
		// if _, ok := value.([]int); ok {
		// 	for _, val := range value.([]int) {
		// 		if val == valInput {
		// 			return true
		// 		}
		// 	}
		// }

	}
	return false
}

func Reduce(valRef *float64, list interface{}) {
	listValue := reflect.ValueOf(list)
	listType := listValue.Type()

	var tmp interface{}

	switch listType.Kind() {
	case reflect.Slice:
		for i := 0; i < listValue.Len(); i++ {
			tmp += listValue.Index(i).Interface()
		}

	case reflect.Array:
		for _, value := range list {
			tmp += value
		}
	}
	for _, value := range list {
		tmp += value
	}
	*valRef = float64(tmp)
}

func MapValInt(m map[interface{}]int) interface{} {
	var tmp int
	for _, val := range m {
		tmp += val
	}
	return tmp
}

func IntersectionInt(sl ...[]int) []int {
	slideOutput := []int{}
	slide1 := sl[0]
	for _, val1 := range slide1 {
		flgInter := true
		for j := 1; j < len(sl); j++ {
			if !Contains(val1, sl[j]) {
				flgInter = false
				continue
			}
		}
		if flgInter {
			slideOutput = append(slideOutput, val1)
		}
	}
	return slideOutput
}

func main() {
	// val1 := 3
	// list1 := []int{1, 2, 3, 4, 5}
	// result := Contains(val1, list1)

	// val1 := "3"
	// list1 := [5]string{"1", "2", "3", "4", "5"}
	// result := Contains(val1, list1)

	// fmt.Println(result)

	// var reslt3 float64
	// list3 := []int{1, 2, 3, 4, 5}
	// Reduce(&reslt3, list3)
	// fmt.Println(reslt3)

	// m := map[interface{}]int{"key1": 1, "key2": 2, "key3": 3}
	// total := MapValInt(m)
	// fmt.Println(total)

	list1 := []int{1, 2, 3, 6}
	list2 := []int{1, 2, 3, 4, 6}
	list3 := []int{3, 2, 6}
	slInter := IntersectionInt(list1, list2, list3)
	fmt.Println(slInter)
}
