package main

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// BouKe model
type BouKe struct {
}

// GetBloggerName get blogger name
func (receiver *BouKe) GetBloggerName() string {
	return "bou.ke"
}

// GetAllPageURLs get all page urls
func (receiver *BouKe) GetAllPageURLs() []string {
	return []string{receiver.GetFirstPageURL()}
}

// GetFirstPageURL get first page url
func (receiver *BouKe) GetFirstPageURL() string {
	return "https://bou.ke/"
}

// GetLinksFromPage get links from page
func (receiver *BouKe) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//div[@class="content"]/ul[@class="posts"]/li/a/@href`, `https://bou.ke`)
}

// GetBlogInfo get blog info
func (receiver *BouKe) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	info, err := GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//div[@class="content"]/div[@class="post"]/h1`,
		`//div[@class="content"]/div[@class="post"]/h1`,
		`//div[@class="content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//div[@class="content"]/div[@class="post"]/p/span[@class="date"]`, "Jan 2006")
		})
	if err != nil {
		return nil, err
	}
	info.Author = "bou.ke"
	return info, nil
}
