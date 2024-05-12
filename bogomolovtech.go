package main

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// BogomolovTech struct
type BogomolovTech struct {
}

// GetBloggerName get blogger name
func (receiver *BogomolovTech) GetBloggerName() string {
	return "bogomolov.tech"
}

// GetAllPageURLs get all page urls
func (receiver *BogomolovTech) GetAllPageURLs() []string {
	return []string{"https://bogomolov.tech/archives/"}
}

// GetFirstPageURL get first page url
func (receiver *BogomolovTech) GetFirstPageURL() string {
	return "https://bogomolov.tech/archives/"
}

// GetLinksFromPage get blog links from page
func (receiver *BogomolovTech) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//li[@class="post-preview"]//div[@class="post-info"]/a/@href`, "https://bogomolov.tech")
}

// GetBlogInfo get blog detail information
func (receiver *BogomolovTech) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="post-title"]`,
		`//div[@class="name"]`,
		`//div[@class="post-content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//div[@class="post-meta"]//span[@class="attr"]/span`, "2006-01-02")
		})
}
