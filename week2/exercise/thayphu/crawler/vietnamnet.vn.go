package crawler

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type impVietNamNet struct {
	Crawler
}

var extractPublishDateVietNamNet = func(selector string, doc *goquery.Document) time.Time {
	publishedDateStr := strings.TrimSpace(extract(selector, doc))
	r, _ := regexp.Compile("[^0-9/:+GMT]+")
	publishedDateStr = string(r.ReplaceAll([]byte(publishedDateStr), []byte(" ")))
	// Value: 24/03/2019 07:50 GMT+7
	// Format: Mon Jan 2 15:04:05 MST 2006 | 2006-01-02T15:04:05.000Z
	publishedDate, _ := time.Parse("02/01/2006 15:04 MST", publishedDateStr)
	return publishedDate
}

func extractAuthorVietNamNet(selector string, doc *goquery.Document) string {
	if selector == "" {
		return ""
	}
	value := doc.Find(selector).Last().Text()
	return strings.TrimSpace(value)
}

func CreateVietNamNetCrawler() ICrawler {
	// Selector de extract tu html download ve dc
	selector := Selector{
		Title:         "title",
		PublishedDate: "#ArticleHolder .ArticleDate.right",
		Author:        "#ArticleContent p>span.bold",
		Content:       "#ArticleContent",
	}
	crawler := &impVietNamNet{}
	crawler.selector = selector
	crawler.parser = createDefaultParser()
	crawler.parser.extractPublishDate = extractPublishDateVietNamNet
	crawler.parser.extractAuthor = extractAuthorVietNamNet
	return crawler
}
