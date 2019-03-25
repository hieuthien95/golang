package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"nordiccoder/week2/exercise/crawler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db               *gorm.DB
	chanUrls         chan Url
	chanSaveActicles chan Article
	chanUrlsUpdate   chan Url
)

type (
	Url struct {
		Id     int
		Url    string
		Status int
	}

	Article struct {
		crawler.Data
		Id int
	}
)

var SaiGonTime crawler.ICrawler = crawler.CreateSaiGonTimeCrawler()

// var VietNamNet crawler.ICrawler = crawler.CreateVietNamNetCrawler()

func main() {
	db, _ = gorm.Open("mysql", "root:123456789@/golang?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(Url{}, Article{})

	chanUrls = make(chan Url)
	chanUrlsUpdate = make(chan Url)
	chanSaveActicles = make(chan Article)

	go load()
	go crawl()
	go save()
	go update()

	// url := "https://vietnamnet.vn/vn/cong-nghe/ung-dung/cach-su-dung-google-maps-de-giam-sat-vi-tri-cua-tre-nho-514378.html"
	// resp, _ := http.Get(url)
	// data := VietNamNet.Parse(resp)
	// b, _ := json.Marshal(data)
	// fmt.Println(string(b))

	time.Sleep(time.Second * 20)
}

func load() {
	urls := []Url{}

	// if err := db.Find(&urls).Error; err != nil {
	if err := db.Find(&urls, "status <> 1").Error; err != nil {
		panic("db.Find(&urls).Error")
	}

	for _, url := range urls {
		chanUrls <- url
	}
}

func crawl() {
	for {
		url := <-chanUrls
		chanUrlsUpdate <- url

		resp, err := http.Get(url.Url)
		if err == nil {
			defer resp.Body.Close()
		}

		data := SaiGonTime.Parse(resp)
		b, _ := json.Marshal(data.Title)
		fmt.Println(string(b))

		article := Article{}
		article.Title = data.Title
		article.PublishedDate = data.PublishedDate
		article.Author = data.Author
		article.Content = data.Content

		chanSaveActicles <- article
	}
}

func save() {
	articles := []Article{}

	for {
		articles = append(articles, <-chanSaveActicles)

		if len(articles) >= 5 {
			for _, atc := range articles {
				atc.Id = 0
				err := db.Save(&atc).Error
				if err != nil {
					panic(err)
				}
			}
			articles = []Article{}
		}
	}
}

func update() {
	for {
		url := <-chanUrlsUpdate

		url.Status = 1
		err := db.Save(&url).Error

		if err != nil {
			panic("db.Save(&url).Error")
		}
	}
}
