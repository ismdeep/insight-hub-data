package main

import (
	"fmt"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// MeituanTech struct
type MeituanTech struct {
}

func (receiver *MeituanTech) Homepage() string {
	return "https://tech.meituan.com/"
}

// GetBloggerName get blogger name
func (receiver *MeituanTech) GetBloggerName() string {
	return "美团技术团队"
}

// GetSourceName get source name
func (receiver *MeituanTech) GetSourceName() string {
	return "tech.meituan.com"
}

// GetAllPageURLs get all page urls
func (receiver *MeituanTech) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	for i := 2; i <= 26; i++ {
		links = append(links, fmt.Sprintf("https://tech.meituan.com/page/%v.html", i))
	}
	return links
}

// GetFirstPageURL get first page url
func (receiver *MeituanTech) GetFirstPageURL() string {
	return "https://tech.meituan.com/"
}

// GetLinksFromPage get blog links from page
func (receiver *MeituanTech) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//h2[@class="post-title"]/a/@href`,
		``)
}

// GetBlogInfo get blog detail information
func (receiver *MeituanTech) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@name="author"]/@content`,
		`//div[@class="post-content"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
