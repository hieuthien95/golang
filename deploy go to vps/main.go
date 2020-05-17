package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type Health struct {
	Status     string
	StatusCode int
	Msg        string
}

func main() {

	e := echo.New()

	e.GET("/healthcheck", func(c echo.Context) error {
		var heal Health
		heal.Msg = "BHT Service Runing"
		heal.StatusCode = http.StatusOK
		heal.Status = "OK"
		return c.JSON(http.StatusOK, heal)
	})

	e.GET("/healthz", func(c echo.Context) error {
		var heal Health
		heal.Msg = "BHT Service Runing"
		heal.StatusCode = http.StatusOK
		heal.Status = "OK"
		return c.JSON(http.StatusOK, heal)
	})

	e.GET("/readyz", func(c echo.Context) error {
		var heal Health
		heal.Msg = "BHT Service Ready"
		heal.StatusCode = http.StatusOK
		heal.Status = "OK"
		return c.JSON(http.StatusOK, heal)
	})

	e.GET("/service", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "service Golang")
	})

	e.GET("/checkHealthGroup", func(c echo.Context) error {

		url := "http://api.thuannc.tech/healthz"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}

		health := new(Health)

		err = json.Unmarshal(body, &health)
		if err != nil {
			log.Println(err)
		}

		return c.JSON(http.StatusOK, health)
	})

	e.GET("/sum", func(c echo.Context) error {

		a, erra := strconv.Atoi(c.QueryParam("input1"))
		b, errb := strconv.Atoi(c.QueryParam("input2"))

		if erra != nil && errb != nil {
			bs := bson.M{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("input1, input2 are not a number"),
				"data": []bson.M{
					bson.M{
						"input1": c.QueryParam("input1"),
						"input2": c.QueryParam("input2"),
						"sum":    "NAN",
					},
				},
			}
			return c.JSON(http.StatusOK, bs)
		}

		bs := bson.M{
			"status":  http.StatusOK,
			"message": "",
			"data": []bson.M{
				bson.M{
					"input1": a,
					"input2": b,
					"sum":    a + b},
			},
		}
		return c.JSON(http.StatusOK, bs)
	})

	err := e.Start(":9092")
	if err != nil {
		fmt.Println(err.Error())
	}
}
