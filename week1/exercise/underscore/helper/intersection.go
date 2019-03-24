package helper

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
