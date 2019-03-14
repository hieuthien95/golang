package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var ()

func getListUser(ctx echo.Context) error {
	var usrs []User

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydemo").C("user")
	c.Find(bson.M{}).All(&usrs)

	return ctx.JSON(http.StatusOK, usrs)
}

func getUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydemo").C("user")
	usr := User{}
	c.Find(bson.M{"id": id}).One(&usr)

	return ctx.JSON(http.StatusOK, usr)
}

func createUser(ctx echo.Context) error {
	var usr = &User{}
	ctx.Bind(usr)

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydemo").C("user")

	c.UpsertId(usr.Name, usr)

	return ctx.JSON(http.StatusCreated, usr)
}

func updateUser(ctx echo.Context) error {
	// var usr = new(User)

	// id, _ := strconv.Atoi(ctx.Param("id"))
	// if err := db.Where("id = ?", id).First(&usr).Error; err != nil {
	// 	ctx.String(http.StatusNotFound, "StatusNotFound!")
	// 	fmt.Println(err)
	// }
	// ctx.Bind(&usr)

	// db.Save(&usr)

	return ctx.JSON(http.StatusOK, nil)
}

func deleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("mydemo").C("user")

	c.RemoveId(id)

	return ctx.NoContent(http.StatusNoContent)
}

func main() {

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
