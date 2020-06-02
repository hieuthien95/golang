package check

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

	for _, check := range arrPayroll {
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
				datas := check.env + "\n" + resp.Message + "\n\n"
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
