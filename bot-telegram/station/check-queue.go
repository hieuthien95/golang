package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	token  string
	offset int
	// +7
	tz = time.FixedZone("UTC+7", +7*60*60)
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

// ========================================================================================================================

type QueueResp struct {
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

func callRestStation(url string, authen string) (resp QueueResp, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resp, err
	}

	req.Header.Add("Authorization", authen)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func getUpdateActionStation() {

	arr := []Env{}

	for _, check := range arr {
		for _, item := range check.items {

			for _, qType := range item.queueTypes {
				url := item.baseURL + "?queueType=" + qType + "&limit=1"
				infoMessage := check.env + " - " + item.system + " - " + qType + " - " + fmt.Sprint(time.Now().In(tz))

				resp, err := callRestStation(url, check.token)
				if err != nil {
					// send message
					sendMessage(-410940764, infoMessage)
					sendMessage(-410940764, err.Error())

					// console
					fmt.Println(infoMessage)
					fmt.Println(err.Error())

					fmt.Println()

				} else if len(resp.Data) > 0 && len(resp.Data[0].Log) >= 5 {

					// send message
					sendMessage(-410940764, infoMessage)
					sendMessage(-410940764, resp.Data[0].Log[0])

					// console
					fmt.Println(infoMessage)
					byte, _ := json.Marshal(resp)
					fmt.Println(string(byte))

					fmt.Println()
				}
			}
		}
	}

}

// ========================================================================================================================

type PayrollResp struct {
	Status  string   `json:"status,omitempty"`
	Data    []string `json:"data,omitempty"`
	Message string   `json:"message,omitempty"`
}

func callRestPayroll(url string, authen string) (data PayrollResp, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	req.Header.Add("Authorization", authen)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getUpdateActionPayroll() {

	arr := []Env{}

	for _, check := range arr {
		for _, item := range check.items {

			url := item.baseURL

			resp, err := callRestPayroll(url, check.token)
			if err != nil {
				// send message
				sendMessage(-410940764, err.Error())

				// console
				fmt.Println(err.Error())

				fmt.Println()

			} else {
				datas := check.env + "\n"
				for _, s := range resp.Data {
					datas += s + "\n"
				}

				// send message
				sendMessage(-410940764, datas)

				// console
				fmt.Println(resp.Data)

				fmt.Println()
			}
		}
	}

}

// ========================================================================================================================

type LichPhatSongResp struct {
	LichPhatSong struct {
		ChannelID int    `json:"ChannelId"`
		Date      string `json:"Date"`
		EventList []struct {
			StartTime          string `json:"StartTime"`
			EndTime            string `json:"EndTime"`
			Name               string `json:"Name"`
			ShortDescriptor    string `json:"ShortDescriptor"`
			ExtendedDescriptor string `json:"ExtendedDescriptor"`
			Hot                bool   `json:"Hot"`
		} `json:"EventList"`
	} `json:"LichPhatSong"`
	End bool   `json:"End"`
	Ext string `json:"Ext"`
}

type Kenh struct {
	Ma  int    `json:"ma"`
	Ten string `json:"ten"`
}

func callRestChannel(url string, bodyData map[string]interface{}) (data LichPhatSongResp, err error) {

	byte, err := json.Marshal(bodyData)
	if err != nil {
		return data, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byte))
	if err != nil {
		return data, err
	}

	// req.Header.Add("Authorization", authen)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getUpdateActionChannel(kenh string) {
	now := time.Now()
	hnow := now.Format("15:04:05")

	kenhs := []Kenh{
		{Ma: 94, Ten: "Animal Planet"},
		{Ma: 50, Ten: "AXN"},
		{Ma: 90, Ten: "Cartoon Network"},
		{Ma: 250, Ten: "Cinemax"},
		{Ma: 92, Ten: "Discovery Channel"},
		{Ma: 26, Ten: "FOX MOVIES"},
		{Ma: 23, Ten: "HBO"},
	}

	kenh = strings.Replace(kenh, "/chan", "", 1)
	kenh = strings.ToLower(kenh)

	for _, k := range kenhs {
		k.Ten = strings.ToLower(k.Ten)
		if strings.Contains(k.Ten, kenh) {

			body := make(map[string]interface{})
			body["maKenh"] = k.Ma
			body["ngay"] = now

			URL := "https://www.sctv.com.vn/WebMain/LichPhatSong/LayLichPhatSong"

			resp, err := callRestChannel(URL, body)
			if err != nil {
				// send message
				sendMessage(-427411096, err.Error())

				// console
				fmt.Println(err.Error())

				fmt.Println()

			} else {
				datas := strings.ToUpper(k.Ten) + "\n ============================= \n\n"
				for _, l := range resp.LichPhatSong.EventList {
					if hnow >= l.EndTime {
						datas += l.StartTime + " - " + l.Name + "\n\n"
					}
				}

				// byte, _ := json.Marshal(resp.LichPhatSong.EventList)
				// fmt.Println(string(byte))

				// send message
				sendMessage(-427411096, fmt.Sprint(datas))

				// console
				fmt.Println(datas)

				fmt.Println()
			}
		}
	}

}

// ========================================================================================================================

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
