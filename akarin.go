package main

import (
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// Akarin struct
type Akarin struct {
}

// GetBloggerName get blogger name
func (receiver *Akarin) GetBloggerName() string {
	return "akarin.dev"
}

// GetAllPageURLs get all page urls
func (receiver *Akarin) GetAllPageURLs() []string {
	return []string{
		receiver.GetFirstPageURL(),
		"https://akarin.dev/page/2/",
		"https://akarin.dev/page/3/",
		"https://akarin.dev/page/4/",
	}
}

// GetFirstPageURL get first page url
func (receiver *Akarin) GetFirstPageURL() string {
	return "https://akarin.dev/"
}

// GetLinksFromPage get blog links from page
func (receiver *Akarin) GetLinksFromPage(pageURL string) ([]string, error) {
	links, err := GetBlogLinks(pageURL,
		`//main[@class="mdui-container"]//div[@class="mdui-typo mdui-card-content"]/a/@href`,
		``)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(links); i++ {
		links[i] = strings.ReplaceAll(links[i], "../../", "")
		links[i] = "https://akarin.dev/" + links[i]
	}

	return links, nil
}

// GetBlogInfo get blog detail information
func (receiver *Akarin) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//meta[@property="og:title"]/@content`,
		`//meta[@property="article:author"]/@content`,
		`//article[@class="mdui-p-a-3 mdui-card-content mdui-typo"]`,
		func(doc *html.Node) (time.Time, error) {
			return GetTime(doc, `//meta[@property="article:published_time"]/@content`, time.RFC3339)
		})
}
