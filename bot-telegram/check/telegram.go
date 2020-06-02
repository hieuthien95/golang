package check

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func sendMessage(chatID int, msg string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v", token, chatID, url.QueryEscape(msg))

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
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update UpdateMessage

	err = json.Unmarshal(body, &update)
	if err != nil {
		log.Println(err)
		return
	}

	for _, elm := range update.Result {
		offset = elm.UpdateID + 1

		// getUpdateActionStation
		if elm.Message.Text == "/station" {
			getUpdateActionStation()
			sendMessage(-410940764, "Done")
			return
		}
		// getUpdateActionPayroll
		if elm.Message.Text == "/payroll" {
			getUpdateActionPayroll()
			sendMessage(-410940764, "Done")
			return
		}
		// channel
		if strings.Contains(elm.Message.Text, "/chan") {
			getUpdateActionChannel(elm.Message.Text)
			sendMessage(-427411096, "Done")
			return
		}

		// channel
		if strings.Contains(elm.Message.Text, "/health") {
			sendMessage(elm.Message.From.ID, elm.Message.Text+" ?")
			return
		}

		fmt.Println(offset, elm.Message.From.Username, elm.Message.Text)
	}

	// fmt.Println(res)
	// fmt.Println(string(body))
}

// ========================================================================================================================
