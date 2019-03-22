package helper

import (
	"reflect"
	"strings"
)

func Contains1(valInput interface{}, arrInput ...interface{}) bool {
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

func Contains2(list interface{}, v interface{}) bool {
	// 1. Can dung reflect de convert interface truyen vao
	// thanh doi tuong reflect.Value
	listValue := reflect.ValueOf(list)
	value := reflect.ValueOf(v)
	// 2. Can dung reflect de convert interface truyen vao
	// thanh doi tuong reflect.Type
	listType := listValue.Type()
	// 3. Ban dau ket qua dc gan la false
	// Muc dich cho nay la de viet Unit Test
	result := false
	// 4. Voi reflect.Type chung ta dung method Kind de biet
	// kieu list la gi
	switch listType.Kind() {
	case reflect.Slice: //4.1 Slice or Array
		// Vi kieu du lieu la slice nen dung ham .Len o day dc
		for i := 0; i < listValue.Len(); i++ {
			// 4.2 Lay gia tri tai vi tri thu i (reflect.Type)
			// Sau do get ra kieu interface{} de compare voi gia tri truyen vao
			if reflect.DeepEqual(listValue.Index(i).Interface(), v) {
				result = true
				break
			}
		}
		break
	case reflect.String: //4.2 String
		result = strings.Contains(listValue.String(), value.String())
		break
	}
	// 5. Tra ve ket qua
	return result
}
