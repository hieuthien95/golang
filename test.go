package main

import (
	"fmt"
	"reflect"
)

func Map(list interface{}, iterateeFunc interface{}) interface{} {
	listValue := reflect.ValueOf(list)
	listType := listValue.Type()

	iterateeFuncValue := reflect.ValueOf(iterateeFunc)
	typeOfResult := reflect.SliceOf(iterateeFuncValue.Type().Out(0))
	result := reflect.MakeSlice(typeOfResult, 0, 0)

	listKind := listType.Kind()
	if listKind == reflect.Slice || listKind == reflect.Array {
		for i := 0; i < listValue.Len(); i++ {
			elem := listValue.Index(i)
			in := []reflect.Value{elem}
			out := iterateeFuncValue.Call(in)[0]
			result = reflect.Append(result, out)
		}

	}
	return result.Interface()
}
func Map2(list1 interface{}, list2 interface{}) interface{} {
	listValue1 := reflect.ValueOf(list1)
	listType1 := listValue1.Type()
	listValue2 := reflect.ValueOf(list2)
	// listType2 := listValue2.Type()

	iterateeFuncValue := reflect.ValueOf(listValue1.Index(0))
	typeOfResult := reflect.SliceOf(iterateeFuncValue.Type())
	result := reflect.MakeSlice(typeOfResult, 0, 0)

	listKind := listType1.Kind()
	if listKind == reflect.Slice {
		for i := 0; i < listValue2.Len(); i++ {
			elem := listValue2.Index(i)
			result = reflect.Append(result, elem)
		}

	}
	return result.Interface()
}

func main() {

	mapfull := Map2([]int{1, 2, 3}, []int{4, 5, 6})

	fmt.Println(mapfull)
}
