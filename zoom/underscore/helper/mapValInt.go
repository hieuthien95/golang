package helper

func MapValInt(m map[interface{}]int) interface{} {
	var tmp int
	for _, val := range m {
		tmp += val
	}
	return tmp
}
