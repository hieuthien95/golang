package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	token string
)

type queue struct {
	Status string `json:"status"`
	Data   []struct {
		ID              string    `json:"_id"`
		ConsumerVersion string    `json:"consumer_version"`
		CreatedTime     time.Time `json:"created_time"`
		Data            struct {
			Data interface{} `json:"data"`
		} `json:"data"`
		LastUpdatedTime time.Time `json:"last_updated_time"`
		ProcessBy       string    `json:"process_by"`
		Log             []string  `json:"log"`
	} `json:"data"`
	Message string `json:"message"`
	Total   int    `json:"total"`
}

func sendMessage(chatID int, msg string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v", token, chatID, msg)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "54526234-aa86-47c7-814f-702d8d3523f3")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

// Check ...
type Check struct {
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

func getUpdate() {

	arr := []Check{
	}

	for _, check := range arr {
		for _, item := range check.items {

			for _, qType := range item.queueTypes {
				url := item.baseURL + "?queueType=" + qType + "&getTotal=true&limit=1"

				queue := checkLogs(url, check.token)
				if len(queue.Data) > 0 && len(queue.Data[0].Log) >= 5 {
					message := check.env + " - " + item.system + " - " + qType
					content := queue.Data[0].Log[0]

					// send message
					sendMessage(702464361, message)
					sendMessage(702464361, content)

					// console
					fmt.Println(message)

					byte, _ := json.Marshal(queue)
					fmt.Println(string(byte))

					fmt.Println()
				}
			}
		}
	}

}

func checkLogs(url string, authen string) queue {
	var queue queue

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return queue
	}

	req.Header.Add("Authorization", authen)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return queue
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
		return queue
	}

	err = json.Unmarshal(body, &queue)
	if err != nil {
		fmt.Print(err.Error())
		return queue
	}

	return queue
}

func main() {
	token = "880494249:AAHY7N-75FacHJNK2HqefQl96mxf7flEC_c"

	func() {
		for {
			getUpdate()
			time.Sleep(time.Second * 60 * 60)
		}
	}()
}
