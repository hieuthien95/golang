package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	usrs = map[int]*User{}
	seq  = 1
)

func getListUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, usrs)
}

func getUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return ctx.JSON(http.StatusOK, usrs[id])
}

func createUser(ctx echo.Context) error {
	u := &User{}
	if err := ctx.Bind(u); err != nil {
		return err
	}
	fmt.Println(">>>", u)
	usrs[u.ID] = u
	seq++
	return ctx.JSON(http.StatusCreated, u)
}

func updateUser(ctx echo.Context) error {
	u := new(User)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	usrs[id].Name = u.Name
	return ctx.JSON(http.StatusOK, usrs[id])
}

func deleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	delete(usrs, id)
	return ctx.NoContent(http.StatusNoContent)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
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
