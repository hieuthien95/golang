package check

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

func checkRTP() string {
	var message string

	if len(arrStation2) != 1 || len(arrStation2[0].items) != 2 {
		message = "Invalid len"

		// send message
		sendMessage(-410940764, message)

		// console
		fmt.Println(message)
		fmt.Println()

		return message
	}
	check := arrStation2[0]

	item := check.items[0]
	url := item.baseURL
	resp1, err := callRestStation(url, check.token)
	if err != nil {
		message = url + "\n\n" + err.Error()

		// send message
		sendMessage(-410940764, message)

		// console
		fmt.Println(message)
		fmt.Println()

		return message
	}

	item = check.items[1]
	url = item.baseURL
	resp2, err := callRestStation(url, check.token)
	if err != nil {
		message = url + "\n\n" + err.Error()

		// send message
		sendMessage(-410940764, message)

		// console
		fmt.Println(message)
		fmt.Println()

		return message
	}

	message = fmt.Sprintf("Queue(%v) - RTP(%v)", resp1.Total, resp2.Total)
	if resp1.Total < resp2.Total || resp1.Total == 0 || resp2.Total == 0 {
		// send message
		sendMessage(-410940764, message)

		// console
		fmt.Println(message)
		fmt.Println()

		return message
	}

	return message
}

func getUpdateActionStation() {
	var message string

	//
	for _, check := range arrStation {
		for _, item := range check.items {

			for _, qType := range item.queueTypes {
				url := item.baseURL + "?queueType=" + qType + "&limit=100"
				infoMessage := check.env + " - " + item.system + " - " + qType + "\n" + fmt.Sprint(time.Now().In(tz))

				resp, err := callRestStation(url, check.token)
				if err != nil {
					message = infoMessage + "\n\n" + err.Error()

					// send message
					sendMessage(-410940764, message)

					// console
					fmt.Println(message)
					fmt.Println()

				} else if len(resp.Data) > 0 && len(resp.Data[0].Log) >= 5 {
					byte, _ := json.Marshal(resp.Data[0].Data)

					message = infoMessage + "\n" + "len: " + fmt.Sprint(len(resp.Data)) + "\n\n" + resp.Data[0].Log[0] + "\n\n" + string(byte)

					// send message
					sendMessage(-410940764, message)

					// console
					fmt.Println(message)
					fmt.Println()
				}
			}
		}
	}

	//
	checkRTP()
}
