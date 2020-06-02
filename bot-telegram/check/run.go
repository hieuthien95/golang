package check

import (
	"fmt"
	"time"
)

var (
	token  string
	offset int
	// +7
	tz = time.FixedZone("UTC+7", +7*60*60)
)

// Env ...
type Env struct {
	env   string
	token string
	items []Item
}

// Item ...
type Item struct {
	system     string
	baseURL    string
	queueTypes []string
}

// Run ...
func Run() {
	fmt.Println("Running ...")
	token = "880494249:AAHY7N-75FacHJNK2HqefQl96mxf7flEC_c"

	go func() {
		for {
			getUpdateMessage()
			time.Sleep(time.Second * 5)
		}
	}()

	func() {
		for {
			getUpdateActionStation()
			time.Sleep(time.Second * 60 * 10)
		}
	}()
}
