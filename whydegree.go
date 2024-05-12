package main

import (
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// WhyDegree struct
type WhyDegree struct {
}

// GetBloggerName get blogger name
func (receiver *WhyDegree) GetBloggerName() string {
	return "why.degree"
}

// GetAllPageURLs get all page urls
func (receiver *WhyDegree) GetAllPageURLs() []string {
	return []string{"https://why.degree/"}
}

// GetFirstPageURL get first page url
func (receiver *WhyDegree) GetFirstPageURL() string {
	return "https://why.degree/"
}

// GetLinksFromPage get blog links from page
func (receiver *WhyDegree) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//ul[@class="post-list"]/li/h3/p[@class="post-link"]/a/@href`, "https://why.degree")
}

// GetBlogInfo get blog detail information
func (receiver *WhyDegree) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//header[@class="post-header"]/h1[@class="post-title p-name"]`,
		`//div[@class="wrapper"]/a[@class="site-title"]`,
		`//div[@class="post-content e-content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//time[@class="dt-published"]/@datetime`, time.RFC3339)
		})
}
