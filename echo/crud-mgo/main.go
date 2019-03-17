package main

import (
	"log"
	"net/http"

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

func init() {
	ssinit, err := mgo.Dial("localhost:27017")
	// s1.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
	}
	session = ssinit
}
func main() {

	defer session.Close()

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

	// fmt.Println("enter main - connecting to mongo")
	// var usrs []User

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

	// maxWait := time.Duration(5 * time.Second)
	// session, sessionErr := mgo.DialWithTimeout("localhost:27017", maxWait)
	// defer session.Close()
	// if sessionErr == nil {
	// 	session.SetMode(mgo.Monotonic, true)
	// 	coll := session.DB("mydemo").C("user")
	// 	if coll != nil {
	// 		fmt.Println("Got a collection object")

	// 		coll.Find(bson.M{}).All(&usrs)
	// 	}
	// } else { // never gets here
	// 	fmt.Println("Unable to connect to local mongo instance!")
	// }

}
