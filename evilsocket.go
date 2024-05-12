package main

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// Evilsocket blogger
type Evilsocket struct {
}

// GetBloggerName get blogger name
func (receiver *Evilsocket) GetBloggerName() string {
	return "www.evilsocket.net"
}

// GetAllPageURLs get all page url
func (receiver *Evilsocket) GetAllPageURLs() []string {
	var lst []string
	lst = append(lst, receiver.GetFirstPageURL())
	for i := 2; i <= 7; i++ {
		lst = append(lst, fmt.Sprintf("https://www.evilsocket.net/page/%v/", i))
	}

	return lst
}

// GetFirstPageURL get first page url
func (receiver *Evilsocket) GetFirstPageURL() string {
	return "https://www.evilsocket.net/"
}

// GetLinksFromPage get links from page
func (receiver *Evilsocket) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//ul[@class="post-list"]/li[@class="post-item"]/span/a/@href`, "https://www.evilsocket.net")
}

// GetBlogInfo get blog info
func (receiver *Evilsocket) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@property="article:author"]/@content`,
		`//article[@class="post"]/div[@class="content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
