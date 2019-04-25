package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	db *gorm.DB
)

type User struct {
	ID   int    `json:"ids"`
	Name string `json:"names"`
}

func ping(c echo.Context) error {

	var usrs []User
	if err := db.Find(&usrs).Error; err != nil {
		c.String(http.StatusNotFound, "StatusNotFound!")
		fmt.Println(err)
	}

	return c.JSONPretty(http.StatusOK, usrs, " ")
}

func updateF(c echo.Context) error {
	id := c.QueryParam("id")
	user := User{}

	err := c.Bind(&user)

	fmt.Print(id, user, err)

	return c.JSON(http.StatusOK, user)
}

func main() {
	db, err := gorm.Open("mysql", "root:123456789@/golang?charset=utf8&parseTime=True&loc=Local")

	db.AutoMigrate(&User{})
	defer db.Close()

	fmt.Print(err)

	e := echo.New()

	e.GET("/ping", ping)
	e.GET("/upadte", updateF)

	e.Start(":9090")
}
