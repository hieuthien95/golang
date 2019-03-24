package helper

func Reduce(valRef *float64, list []int) {
	var tmp int
	for _, value := range list {
		tmp += value
	}
	*valRef = float64(tmp)
}
