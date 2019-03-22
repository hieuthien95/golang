package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	db *gorm.DB
)

func getListUser(ctx echo.Context) error {
	var usrs []User
	if err := db.Find(&usrs).Error; err != nil {
		ctx.String(http.StatusNotFound, "StatusNotFound!")
		fmt.Println(err)
	}

	return ctx.JSONPretty(http.StatusOK, usrs, " ")
}

func getUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var usr User
	if err := db.Where("id = ?", id).First(&usr).Error; err != nil {
		ctx.String(http.StatusNotFound, "StatusNotFound!")
		fmt.Println(err)
	}

	return ctx.JSON(http.StatusOK, usr)
}

func createUser(ctx echo.Context) error {
	var usr = &User{}

	ctx.Bind(usr)
	db.Create(&usr)

	return ctx.JSON(http.StatusCreated, usr)
}

func updateUser(ctx echo.Context) error {
	var usr = new(User)

	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := db.Where("id = ?", id).First(&usr).Error; err != nil {
		ctx.String(http.StatusNotFound, "StatusNotFound!")
		fmt.Println(err)
	}
	ctx.Bind(&usr)

	db.Save(&usr)

	return ctx.JSON(http.StatusOK, usr)
}

func deleteUser(ctx echo.Context) error {
	var usr User

	id, _ := strconv.Atoi(ctx.Param("id"))
	db.Where("id = ?", id).Delete(&usr)

	return ctx.NoContent(http.StatusNoContent)
}

func main() {
	db, _ = gorm.Open("mysql", "root:123456789@/golang?charset=utf8&parseTime=True&loc=Local")

	db.AutoMigrate(&User{})
	defer db.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getListUser)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
