package main

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// CoolShell struct
type CoolShell struct {
}

// GetBloggerName get blogger name
func (receiver *CoolShell) GetBloggerName() string {
	return "coolshell.cn"
}

// GetAllPageURLs get all page urls
func (receiver *CoolShell) GetAllPageURLs() []string {
	return []string{"https://coolshell.cn/featured"}
}

// GetFirstPageURL get first page url
func (receiver *CoolShell) GetFirstPageURL() string {
	return "https://coolshell.cn/featured"
}

// GetLinksFromPage get blog links from page
func (receiver *CoolShell) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//ul[@class="featured-post"]/li//a/@href`, "")
}

// GetBlogInfo get blog detail information
func (receiver *CoolShell) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//header[@class="entry-header"]/h1[@class="entry-title"]`,
		`//span[@class="author vcard"]/a`,
		`//div[@class="entry-content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//time[@class="entry-date"]/@datetime`, time.RFC3339)
		})
}
