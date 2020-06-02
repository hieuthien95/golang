package check

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

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
	now := time.Now().In(tz)
	hnow := now.Format("15:04:05")

	kenh = strings.Replace(kenh, "/chan", "", 1)
	kenh = strings.ToLower(kenh)

	for _, k := range arrKenh {
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
					if hnow <= l.StartTime || hnow <= l.EndTime {
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
