package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo"
)

type (
	Employee struct {
		ID             string `json:"id"`
		EmployeeName   string `json:"employee_name"`
		EmployeeSalary string `json:"employee_salary"`
		EmployeeAge    string `json:"employee_age"`
		ProfileImage   string `json:"profile_image"`
	}

	User struct {
		Name   string
		Age    int
		Salary int
	}
)

func getEmployees(c echo.Context) error {
	// Build the request
	req, err := http.NewRequest("GET", "http://dummy.restapiexample.com/api/v1/employees", nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the data with the data from the JSON
	var data []Employee

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, data)
}

func main() {
	e := echo.New()

	e.GET("/employees", getEmployees)

	e.GET("/marshal", func(c echo.Context) error {
		data := User{Name: "Rachel", Age: 24, Salary: 344444}
		emp, _ := json.Marshal(data)

		pp.Println(emp)
		fmt.Println(string(emp))

		fmt.Println()
		fmt.Println()
		return c.JSON(http.StatusOK, string(emp))
	})
	e.GET("/unmarshal", func(c echo.Context) error {

		data1 := []byte("{\"Name\":\"Rachel\",\"Age\":24,\"Salary\":344444}")
		data2 := []byte{
			0x7b, 0x22, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x52, 0x61, 0x63, 0x68, 0x65, 0x6c, 0x22,
			0x2c, 0x22, 0x41, 0x67, 0x65, 0x22, 0x3a, 0x32, 0x34, 0x2c, 0x22, 0x53, 0x61, 0x6c, 0x61, 0x72,
			0x79, 0x22, 0x3a, 0x33, 0x34, 0x34, 0x34, 0x34, 0x34, 0x7d,
		}

		var result1 User
		var result2 User

		json.Unmarshal(data1, &result1)
		json.Unmarshal(data2, &result2)

		fmt.Println(result1)
		fmt.Println(result2)

		fmt.Println()
		fmt.Println()
		return c.JSON(http.StatusOK, result2)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
