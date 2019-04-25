package pkg

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
