package helper

import "reflect"

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

type convert1 func(int) int

// func Map2(list []int, fn convert1) []int {
// 	result := make([]int, len(list))
// 	for i := 0; i < len(list); i++ {
// 		result[i] = fn(list[i])
// 	}
// 	return result
// }
// list := []int{1, 2, 3}
func Map2(list interface{}, fn interface{}) interface{} {
	// 1. Slice or Array
	// 1.1 Lay value of interface
	listValue := reflect.ValueOf(list)
	// 1.2 Lay cai
	listType := listValue.Type()
	listKind := listType.Kind()

	if listKind != reflect.Slice && listKind != reflect.Array {
		panic("list should be slice or array")
	}

	// 2
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()
	outType := fnType.Out(0)

	// 3. Tao ra ra retResult
	// 3.1 Tao ra cai sliceType
	outSliceType := reflect.SliceOf(outType)
	retResult := reflect.MakeSlice(outSliceType, 0, 0)

	// 4. Loop qua list
	for i := 0; i < listValue.Len(); i++ {
		item := listValue.Index(i)
		// 4.1 Tao input cho func
		in := []reflect.Value{item}
		// 4.2 Call function
		result := fnValue.Call(in)[0]
		// 4.3 Append vao ket qua tu buoc 3 tao ra
		retResult = reflect.Append(retResult, result)
	}
	// 5. Tra ve du lieu
	return retResult.Interface()
}

// list := []int{1, 2, 3}
// list := []interface{}{1, false, "phu"}
// func Map3(list []interface{}, fn interface{}) interface{} {

// 	return nil
// }
