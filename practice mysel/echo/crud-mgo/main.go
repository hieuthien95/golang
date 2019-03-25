package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address string  `json:"address"`
		Exp     float32 `json:"exp"`
	}
)

var (
	session *mgo.Session
)

func getUser(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	id := ctx.Param("id")
	usr := User{}

	c := ss.DB("mydemo").C("user")
	c.Find(bson.M{"name": id}).One(&usr)

	return ctx.JSON(http.StatusOK, usr)
}

func createUser(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	var usr = &User{}
	ctx.Bind(usr)

	c := ss.DB("mydemo").C("user")
	c.Insert(usr)

	return ctx.JSON(http.StatusCreated, nil)
}

func updateUser(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	id := ctx.Param("id")

	usr := &User{}
	ctx.Bind(usr)

	c := ss.DB("mydemo").C("user")
	c.Upsert(bson.M{"name": id}, usr)

	return ctx.JSON(http.StatusOK, nil)
}

func deleteUser(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	id := ctx.Param("id")
	c := ss.DB("mydemo").C("user")
	c.Remove(bson.M{"name": id})

	return ctx.NoContent(http.StatusNoContent)
}

func getListUser(ctx echo.Context) error {
	// Cach 1:
	ss := session.Clone()
	defer ss.Close()

	var usrs []User
	c := ss.DB("mydemo").C("user")
	c.Find(bson.M{}).All(&usrs)

	// Cach 2: ERROR
	// var usrs []User
	// c := session.DB("mydemo").C("user")
	// c.Find(bson.M{}).All(&usrs)

	return ctx.JSON(http.StatusOK, usrs)
}

func bulkUser(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	var usr = &User{}
	ctx.Bind(usr)

	bulk := ss.DB("mydemo").C("user").Bulk()

	bulk.Insert(usr)
	bulk.Update(usr, bson.M{"$set": bson.M{"name": "User_" + usr.Name}})
	bulk.Remove(bson.M{"name": "Thien"})

	bulk.Run()

	return ctx.JSON(http.StatusCreated, nil)
}

func aggregateUser(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	c := ss.DB("mydemo").C("user")

	pipeline := []bson.M{
		// {"$match": bson.M{"address": "MT"}},
		{"$group": bson.M{
			"_id":   "$address",
			"total": bson.M{"$sum": "$age"},
		},
		},
		{"$sort": bson.M{"age": 1}},
	}
	pipe := c.Pipe(pipeline)

	result := []bson.M{}
	_ = pipe.All(&result)

	return ctx.JSON(http.StatusCreated, result)
}

func register(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	d := ss.DB("mydemo")

	err := d.AddUser("hieuthien95", "123456789", false)

	return ctx.JSON(http.StatusCreated, err)
}

func login(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	d := ss.DB("mydemo")

	err := d.Login("hieuthien95", "123456789")

	return ctx.JSON(http.StatusCreated, err)
}

func logout(ctx echo.Context) error {
	ss := session.Clone()
	defer ss.Close()

	d := ss.DB("mydemo")

	d.Logout()

	return ctx.JSON(http.StatusCreated, nil)
}

func init() {
	// ssinit, err := mgo.Dial("localhost:27017")
	ssinit, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "mydemo",
		Username: "root",
		Password: "123456789",
		Timeout:  60 * time.Second,
	})
	// s1.SetMode(mgo.Monotonic, true)
	if err != nil || ssinit == nil {
		panic(fmt.Sprint("Could not connect to mongo: \n\n", err.Error()))
	}
	session = ssinit
}

func main() {

	defer session.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
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

	e.POST("/users/bulk", bulkUser)
	e.GET("/users/aggregate", aggregateUser)

	e.POST("/register", register)
	e.POST("/login", login)
	e.GET("/logout", logout)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))

	// // tried doing this - doesn't work as intended
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Detected panic")
	// 		var ok bool
	// 		err, ok := r.(error)
	// 		if !ok {
	// 			fmt.Printf("pkg:  %v,  error: %s", r, err)
	// 		}
	// 	}
	// }()

}
