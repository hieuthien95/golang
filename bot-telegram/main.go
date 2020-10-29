package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	token  string
	offset int
)

type Updates struct {
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

type YesNoWtf struct {
	Answer string `json:"answer"`
	Forced bool   `json:"forced"`
	Image  string `json:"image"`
}

func getMe() {

	url := fmt.Sprintf("https://api.telegram.org/bot%v/getMe", token)

	req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("Postman-Token", "aa8f27fb-a815-4545-8531-f0230a0d981b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func sendMessage(chatId int, msg string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?chat_id=%v&text=%v", token, chatId, msg)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "54526234-aa86-47c7-814f-702d8d3523f3")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func sendPhoto(chatId int, imgUrl string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendPhoto?chat_id=%v&photo=%v", token, chatId, imgUrl)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "0e1e6bff-1456-4b2f-9b32-e370f0a1af91")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getImageYesNoWtf() string {
	url := "https://yesno.wtf/api"

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

	yesno := new(YesNoWtf)

	err = json.Unmarshal(body, &yesno)
	if err != nil {
		log.Println(err)
	}

	return yesno.Image
}

func getUpdate() {
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

	update := new(Updates)

	err = json.Unmarshal(body, &update)
	if err != nil {
		log.Println(err)
	}

	for _, elm := range update.Result {
		offset = elm.UpdateID + 1
		fmt.Println(offset, elm.Message.From.Username, elm.Message.Text)

		sendMessage(elm.Message.From.ID, elm.Message.Text+" ?")

		url := getImageYesNoWtf()
		sendPhoto(elm.Message.From.ID, url)
	}

	// fmt.Println(res)
	// fmt.Println(string(body))
}

func main() {
	token = "880494249:AAHY7N-75FacHJNK2HqefQl96mxf7flEC_c"
	offset = 0

	// getMe()
	func() {
		for {
			getUpdate()
		}
	}()
}
