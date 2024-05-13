package main

import (
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// GoDev struct
type GoDev struct {
}

func (receiver *GoDev) Homepage() string {
	return "https://go.dev/blog/"
}

// GetBloggerName get blogger name
func (receiver *GoDev) GetBloggerName() string {
	return "The Go Blog"
}

// GetSourceName get source name
func (receiver *GoDev) GetSourceName() string {
	return "go.dev"
}

// GetAllPageURLs get all page urls
func (receiver *GoDev) GetAllPageURLs() []string {
	return []string{"https://go.dev/blog/all"}
}

// GetFirstPageURL get first page url
func (receiver *GoDev) GetFirstPageURL() string {
	return "https://go.dev/blog/all"
}

// GetLinksFromPage get blog links from page
func (receiver *GoDev) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL, `//p[@class="blogtitle"]/a/@href`, "https://go.dev")
}

// GetBlogInfo get blog detail information
func (receiver *GoDev) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	v, err := GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//div[@class="Article"]/h1[2]`,
		`//p[@class="author"]`,
		`//div[@class="Article"]`,
		func(doc *html.Node) (time.Time, error) {
			return time.Now(), nil
		})
	if err != nil {
		return nil, err
	}

	s := v.Author
	if strings.Index(s, "\n") > 0 {
		author := s[:strings.Index(s, "\n")]
		s2 := s[strings.Index(s, "\n")+1:]
		s2 = strings.TrimSpace(s2)

		t, err := time.Parse("2 January 2006", s2)
		if err != nil {
			return nil, err
		}

		v.Author = strings.TrimSpace(author)
		v.Date = t
	} else {
		t, err := time.Parse("2 January 2006", strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}
		v.Author = "go.dev"
		v.Date = t
	}

	return v, nil

}
