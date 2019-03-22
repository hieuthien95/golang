package main

import (
	"fmt"
	"strings"
	"time"
)

func DateFormat(input string) interface{} {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("Opps1: ", err)
		return nil
	}
	split := strings.Split(input, ",")
	split[0] = ""

	layout := " 02/1/2006 03:04"
	time, err := time.ParseInLocation(layout, strings.Join(split, ""), location)
	if err != nil {
		fmt.Println("Opps2: ", err)
		return nil
	}
	return time
}

func main() {
	time := DateFormat("Thứ Sáu, 15/3/2019, 11:11")
	fmt.Println(time)
}
