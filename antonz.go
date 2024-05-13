package main

import (
	"errors"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// Antonz model
type Antonz struct {
}

func (receiver *Antonz) Homepage() string {
	return "https://antonz.org/"
}

// GetBloggerName get blogger name
func (receiver *Antonz) GetBloggerName() string {
	return "Anton Zhiyanov"
}

func (receiver *Antonz) GetSourceName() string {
	return "antonz.org"
}

// GetAllPageURLs get all page urls
func (receiver *Antonz) GetAllPageURLs() []string {
	return []string{receiver.GetFirstPageURL()}
}

// GetFirstPageURL get first page url
func (receiver *Antonz) GetFirstPageURL() string {
	return "https://antonz.org/all/"
}

// GetLinksFromPage get links from page
func (receiver *Antonz) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//div[@class="posts"]//div[@class="post-stub"]//a[@class="post-stub__title"]/@href`,
		"https://antonz.org")
}

// GetBlogInfo get blog info
func (receiver *Antonz) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@name="author"]/@content`,
		`//article[@class="post"]`,
		func(doc *html.Node) (time.Time, error) {
			s := htmlquery.FindOne(doc, `//footer[@class="post__footer"]/div[@class="row"]//div[@class="post__date"]/time/@datetime`)
			if s == nil {
				return time.Now(), nil
			}
			ss := htmlquery.InnerText(s)
			ss = strings.TrimSpace(ss)
			t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", ss)
			if err == nil {
				return t, nil
			}

			t2, err := time.Parse("2006-01-02 15:04:05 -0700 -0700", ss)
			if err == nil {
				return t2, nil
			}

			return time.Now(), errors.New("failed to parse time")
		},
	)
}
