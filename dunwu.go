package main

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// DunWu struct
type DunWu struct {
}

func (receiver *DunWu) Homepage() string {
	return "https://dunwu.github.io/blog/"
}

// GetBloggerName get blogger name
func (receiver *DunWu) GetBloggerName() string {
	return "钝悟"
}

// GetSourceName get source name
func (receiver *DunWu) GetSourceName() string {
	return "dunwu"
}

// GetAllPageURLs get all page urls
func (receiver *DunWu) GetAllPageURLs() []string {
	lst := make([]string, 0)
	lst = append(lst, "https://dunwu.github.io/blog/archives/")
	for i := 2; i <= 32; i++ {
		lst = append(lst, fmt.Sprintf("https://dunwu.github.io/blog/archives/page/%v/", i))
	}
	return lst
}

// GetFirstPageURL get first page url
func (receiver *DunWu) GetFirstPageURL() string {
	return "https://dunwu.github.io/blog/archives/"
}

// GetLinksFromPage get blog links from page
func (receiver *DunWu) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="post-content"]//header[@class="post-header"]/div[@class="post-title"]/a[@class="post-title-link"]/@href`,
		`https://dunwu.github.io`)
}

// GetBlogInfo get blog detail information
func (receiver *DunWu) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="post-title"]`,
		`//p[@class="site-author-name"]`,
		`//div[@class="post-body"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
