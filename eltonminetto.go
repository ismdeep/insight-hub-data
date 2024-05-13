package main

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// EltonMinetto struct
type EltonMinetto struct {
}

func (receiver *EltonMinetto) Homepage() string {
	return "https://eltonminetto.dev/"
}

// GetBloggerName get blogger name
func (receiver *EltonMinetto) GetBloggerName() string {
	return "Meu nome é Elton Minetto"
}

// GetSourceName get source name
func (receiver *EltonMinetto) GetSourceName() string {
	return "eltonminetto.dev"
}

// GetAllPageURLs get all page urls
func (receiver *EltonMinetto) GetAllPageURLs() []string {
	return []string{
		"https://eltonminetto.dev/en/",
		"https://eltonminetto.dev/en/page/2/",
		"https://eltonminetto.dev/en/page/3/",
	}
}

// GetFirstPageURL get first page url
func (receiver *EltonMinetto) GetFirstPageURL() string {
	return "https://eltonminetto.dev/en/"
}

// GetLinksFromPage get blog links from page
func (receiver *EltonMinetto) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="blog-list"]/div//h2/a/@href`,
		``)

}

// GetBlogInfo get blog detail information
func (receiver *EltonMinetto) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//div[@class="page-title text-center"]//h1`,
		`//div[@class="container"]//div[@class="s-post-details"]//p/a`,
		`//div[@class="container"]//div[@class="post-text"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
