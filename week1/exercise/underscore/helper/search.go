package helper

import (
	"reflect"
)

func Find(arr interface{}, predicate interface{}) interface{} {

	iteratorValue := reflect.ValueOf(arr)
	funcValue := reflect.ValueOf(predicate)

	for i := 0; i < iteratorValue.Len(); i++ {
		item := iteratorValue.Index(i)
		in := []reflect.Value{item}
		out := funcValue.Call(in)
		result := out[0].Bool()

		if result == true {
			return item.Interface()
		}
	}
	return nil
}

func isFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// Find can return first item matched in array
func Find2(arr interface{}, predicate interface{}) interface{} {
	arrValue := reflect.ValueOf(arr)
	predicateValue := reflect.ValueOf(predicate)

	//
	var res interface{}

	if !isFunc(predicate) {
		for index := 0; index < arrValue.Len(); index++ {
			if predicateValue.Interface() == arrValue.Index(index).Interface() {
				res = arrValue.Index(index).Interface()
				break
			}
		}
	} else {
		for index := 0; index < arrValue.Len(); index++ {
			elem := arrValue.Index(index)
			in := []reflect.Value{elem}
			result := predicateValue.Call(in)[0]
			if result.Bool() == true {
				res = arrValue.Index(index).Interface()
				break
			}
		}
	}

	return res
}
