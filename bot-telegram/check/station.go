package check

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

	for _, check := range arrStation {
		for _, item := range check.items {

			for _, qType := range item.queueTypes {
				url := item.baseURL + "?queueType=" + qType + "&limit=1"
				infoMessage := check.env + " - " + item.system + " - " + qType + "\n" + fmt.Sprint(time.Now().In(tz))

				resp, err := callRestStation(url, check.token)
				if err != nil {
					// send message
					sendMessage(-410940764, infoMessage+"\n\n"+err.Error())

					// console
					fmt.Println(infoMessage)
					fmt.Println(err.Error())

					fmt.Println()

				} else if len(resp.Data) > 0 && len(resp.Data[0].Log) >= 5 {
					byte, _ := json.Marshal(resp.Data[0].Data)

					// send message
					sendMessage(-410940764, infoMessage+"\n\n"+resp.Data[0].Log[0]+"\n\n"+string(byte))

					// console
					fmt.Println(infoMessage)
					fmt.Println(string(byte))

					fmt.Println()
				}
			}
		}
	}

}
