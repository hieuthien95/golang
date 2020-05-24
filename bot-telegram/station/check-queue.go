package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	token  string
	offset int
	// +7
	tz = time.FixedZone("UTC+7", +7*60*60)
)

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

type UpdateMessage struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date     int    `json:"date"`
			Text     string `json:"text"`
			Entities []struct {
				Offset int    `json:"offset"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"entities"`
		} `json:"message,omitempty"`
	} `json:"result"`
}

func getUpdateMessage() {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/getUpdates?timeout=60&offset=%v", token, offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	// req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("Postman-Token", "bd8ff577-ac16-4c12-9ebb-978079499a7f")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var update UpdateMessage

	err = json.Unmarshal(body, &update)
	if err != nil {
		log.Println(err)
	}

	for _, elm := range update.Result {
		offset = elm.UpdateID + 1

		// getUpdateActionStation
		if elm.Message.Text == "/check" {
			getUpdateActionStation()
			sendMessage(-410940764, "Done")
			return
		}
		// getUpdateActionPayroll
		if elm.Message.Text == "/check-payroll" {
			getUpdateActionPayroll()
			sendMessage(-410940764, "Done")
			return
		}

		fmt.Println(offset, elm.Message.From.Username, elm.Message.Text)

		sendMessage(elm.Message.From.ID, elm.Message.Text+" ?")
	}

	// fmt.Println(res)
	// fmt.Println(string(body))
}

type queue struct {
	Status string `json:"status,omitempty"`
	Data   []struct {
		ID              string    `json:"_id,omitempty"`
		ConsumerVersion string    `json:"consumer_version,omitempty"`
		CreatedTime     time.Time `json:"created_time,omitempty"`
		Data            struct {
			Data interface{} `json:"data,omitempty"`
		} `json:"data,omitempty"`
		LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
		ProcessBy       string    `json:"process_by,omitempty"`
		Log             []string  `json:"log,omitempty"`
	} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Total   int    `json:"total,omitempty"`
}

// Station ...
type Station struct {
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

func getUpdateActionStation() {

	arr := []Station{}

	for _, check := range arr {
		for _, item := range check.items {

			for _, qType := range item.queueTypes {
				url := item.baseURL + "?queueType=" + qType + "&limit=1"
				infoMessage := check.env + " - " + item.system + " - " + qType + " - " + fmt.Sprint(time.Now().In(tz))

				queue, err := callRest(url, check.token)
				if err != nil {
					// send message
					sendMessage(-410940764, infoMessage)
					sendMessage(-410940764, err.Error())

					// console
					fmt.Println(infoMessage)
					fmt.Println(err.Error())

					fmt.Println()

				} else if len(queue.Data) > 0 && len(queue.Data[0].Log) >= 5 {
					content := queue.Data[0].Log[0]

					// send message
					sendMessage(-410940764, infoMessage)
					sendMessage(-410940764, content)

					// console
					fmt.Println(infoMessage)

					byte, _ := json.Marshal(queue)
					fmt.Println(string(byte))

					fmt.Println()
				}
			}
		}
	}

}

func getUpdateActionPayroll() {

	arr := []Station{}

	for _, check := range arr {
		for _, item := range check.items {

			url := item.baseURL

			resp, err := callRest(url, check.token)
			if err != nil {
				// send message
				sendMessage(-410940764, err.Error())

				// console
				fmt.Println(err.Error())

				fmt.Println()

			} else {
				byte, _ := json.Marshal(resp.Data)
				fmt.Println(string(byte))

				// send message
				sendMessage(-410940764, string(byte))

				// console
				fmt.Println(string(byte))

				fmt.Println()
			}
		}
	}

}

func callRest(url string, authen string) (queue, error) {
	var queue queue

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return queue, err
	}

	req.Header.Add("Authorization", authen)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return queue, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return queue, err
	}

	err = json.Unmarshal(body, &queue)
	if err != nil {
		return queue, err
	}

	return queue, nil
}

func main() {
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
