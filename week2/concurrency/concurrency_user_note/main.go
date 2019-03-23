package main

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type (
	Note struct {
		gorm.Model
		Id        int
		Name      string
		IdCreater string
	}
	User struct {
		gorm.Model
		Id   int
		Name string
	}
)

func getNote(db *gorm.DB, id int) (*Note, error) {
	note := new(Note)
	err := db.Where("id = ?", id).First(&note).Error

	return note, err
}

func getCreater(db *gorm.DB, id int) (*User, error) {
	user := new(User)
	err := db.Where("id = ?", id).First(&user).Error

	return user, err
}

func init() {
}

func main() {
	db, err := gorm.Open("mysql", "root:123456789@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Note{}, &User{})

	craterId := 1

	note := &Note{}
	user := &User{}

	// note, _ = getNote(db, craterId)
	// user, _ = getCreater(db, craterId)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		note, _ = getNote(db, craterId)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user, _ = getCreater(db, craterId)
	}()

	wg.Wait()

	fmt.Println("note", note)
	fmt.Println("user", user)
}
