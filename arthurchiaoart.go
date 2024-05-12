package main

import (
	"errors"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
)

// ArthurChiaoArt struct
type ArthurChiaoArt struct {
}

// GetBloggerName get blogger name
func (receiver *ArthurChiaoArt) GetBloggerName() string {
	return "arthurchiao.art"
}

// GetAllPageURLs get all page urls
func (receiver *ArthurChiaoArt) GetAllPageURLs() []string {
	links := []string{receiver.GetFirstPageURL()}
	links = append(links, "http://arthurchiao.art/articles/")
	return links
}

// GetFirstPageURL get first page url
func (receiver *ArthurChiaoArt) GetFirstPageURL() string {
	return "http://arthurchiao.art/articles-zh/"
}

// GetLinksFromPage get blog links from page
func (receiver *ArthurChiaoArt) GetLinksFromPage(pageURL string) ([]string, error) {
	return GetBlogLinks(pageURL,
		`//ul[@class="posts noList"]/li/a/@href`,
		`https://arthurchiao.art`)
}

// GetBlogInfo get blog detail information
func (receiver *ArthurChiaoArt) GetBlogInfo(blogLink string) (*schema.Blog, error) {
	return GetBlogInfo(blogLink, receiver.GetBloggerName(),
		`//h1[@class="postTitle"]`,
		`//p[@class="copy"]/a`,
		`//div[@class="post"]`,
		func(doc *html.Node) (time.Time, error) {
			node := htmlquery.FindOne(doc, `//p[@class="meta"]`)
			if node == nil {
				return time.Now(), errors.New("node not found")
			}
			s := htmlquery.InnerText(node)
			s = strings.Split(s, "|")[0]
			s = strings.ReplaceAll(s, "Published at ", "")
			s = strings.TrimSpace(s)
			return time.Parse("2006-01-02", s)
		})
}
