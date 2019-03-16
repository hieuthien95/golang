package helper

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
