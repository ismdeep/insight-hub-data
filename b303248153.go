package main

import (
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// B303248153 model
type B303248153 struct {
}

func (receiver *B303248153) Homepage() string {
	return "https://303248153.github.io/"
}

// GetBloggerName get blogger name
func (receiver *B303248153) GetBloggerName() string {
	return "老农的博客"
}

// GetSourceName get source name
func (receiver *B303248153) GetSourceName() string {
	return "303248153.github.io"
}

// GetAllPageURLs get all page urls
func (receiver *B303248153) GetAllPageURLs() []string {
	return []string{receiver.GetFirstPageURL()}
}

// GetFirstPageURL get first page url
func (receiver *B303248153) GetFirstPageURL() string {
	return "https://303248153.github.io/"
}

// GetLinksFromPage get links from page
func (receiver *B303248153) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//div[@class="posts"]/ul/li/a/@href`, "https://303248153.github.io")
}

// GetBlogInfo get blog info
func (receiver *B303248153) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@name="author"]/@content`,
		`//article[@class="post"]`,
		func(doc *html.Node) (time.Time, error) {
			s := htmlquery.FindOne(doc, `//div[@class="date"]`)
			if s == nil {
				return time.Now(), nil
			}
			ss := htmlquery.InnerText(s)
			ss = strings.TrimSpace(ss)
			ss = strings.ReplaceAll(ss, "Written on ", "")
			return time.Parse("January 02, 2006", ss)
		})
}
