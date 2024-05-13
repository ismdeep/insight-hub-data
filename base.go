package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"

	"github.com/ismdeep/insight-hub-data/internal/schema"
	"github.com/ismdeep/insight-hub-data/pkg/insight-hub-data/core"
)

// BloggerInterface blogger interface
type BloggerInterface interface {
	GetBloggerName() string
	GetSourceName() string
	GetAllPageURLs() []string
	GetFirstPageURL() string
	GetLinksFromPage(pageURL string) ([]string, error)
	GetBlogInfo(blogLink string) (*schema.Blog, error)
	Homepage() string
}

// GetHTML get html
func GetHTML(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

// GetHTMLDoc get html doc
func GetHTMLDoc(url string) (*html.Node, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3776.0 Safari/537.36")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// GetBlogLinks get blog links
func GetBlogLinks(url string, listXPath string, hrefPrefix string) ([]string, error) {
	doc, err := GetHTMLDoc(url)
	if err != nil {
		return nil, err
	}
	list := htmlquery.Find(doc, listXPath)
	var links []string
	for _, node := range list {
		if node == nil {
			continue
		}
		href := htmlquery.InnerText(node)
		href = strings.TrimSpace(href)
		links = append(links, hrefPrefix+href)
	}
	return links, nil
}

// GetBlogInfo get blog detail information
func GetBlogInfo(url string, source string, titleXPath string, authorXPath string, contentXPath string, parseTime func(doc *html.Node) (time.Time, error)) (*schema.Blog, error) {
	doc, err := GetHTMLDoc(url)
	if err != nil {
		return nil, err
	}

	t, err := parseTime(doc)
	if err != nil {
		return nil, err
	}

	titleNode := htmlquery.FindOne(doc, titleXPath)
	if titleNode == nil {
		return nil, errors.New("title not found")
	}
	title := strings.TrimSpace(htmlquery.InnerText(titleNode))
	if title == "" {
		return nil, errors.New("bad url")
	}

	authorNode := htmlquery.FindOne(doc, authorXPath)
	if authorNode == nil {
		return nil, errors.New("author not found")
	}
	author := strings.TrimSpace(htmlquery.InnerText(authorNode))
	if author == "" {
		return nil, errors.New("bad url")
	}

	contentNode := htmlquery.FindOne(doc, contentXPath)
	if contentNode == nil {
		return nil, errors.New("content not found")
	}
	content := htmlquery.OutputHTML(contentNode, true)

	return &schema.Blog{
		Source:  source,
		Title:   title,
		Link:    url,
		Content: content,
		Date:    t,
		Author:  author,
	}, nil
}

// GetTime get time
func GetTime(top *html.Node, expr string, timeFormat string) (time.Time, error) {
	node := htmlquery.FindOne(top, expr)
	if node == nil {
		return time.Now(), errors.New("time not found")
	}

	return time.Parse(timeFormat,
		htmlquery.InnerText(node))
}

func Download(blogger BloggerInterface) error {

	source := blogger.GetSourceName()
	indexFile := fmt.Sprintf("data/%v.txt", source)

	f, err := os.OpenFile(indexFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)
	if err != nil {
		return err
	}

	metaDataWriter, err := os.OpenFile(fmt.Sprintf("data/%v.meta.json", source), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
	if err != nil {
		return err
	}

	s := core.NewStore(f, metaDataWriter, func(r core.Record) error {
		r.ID = core.RecordID(r)
		raw, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			return err
		}

		if err := os.MkdirAll(fmt.Sprintf("data/%v.d/", source), 0755); err != nil {
			fmt.Println("failed to create directory:", err.Error())
			return err
		}

		if err := os.WriteFile(fmt.Sprintf("data/%v.d/%v.json", source, r.ID), raw, 0644); err != nil {
			fmt.Println("failed to write file:", err.Error())
			return err
		}

		return nil
	})

	raw, err := os.ReadFile(indexFile)
	if err != nil {
		return err
	}
	if err := s.Load(bytes.NewBuffer(raw)); err != nil {
		return err
	}

	if err := s.WriteMeta(core.Meta{
		Source:   blogger.GetSourceName(),
		HomePage: blogger.Homepage(),
		Name:     blogger.GetBloggerName(),
	}); err != nil {
		return err
	}

	pageURLs := blogger.GetAllPageURLs()
	for _, l := range pageURLs {
		links, err := blogger.GetLinksFromPage(l)
		if err != nil {
			fmt.Printf("failed to get links from page: %v\n", err.Error())
			continue
		}
		for _, link := range links {
			if !core.LinkIsTidy(link) {
				fmt.Printf("link is not tidy: %v\n", link)
				continue
			}

			if s.URLExists(link) {
				continue
			}

			info, err := blogger.GetBlogInfo(link)
			if err != nil {
				fmt.Printf("failed to get blog info %v: %v\n", link, err.Error())
				continue
			}
			if err := s.Save(core.Record{
				Source:      info.Source,
				Link:        info.Link,
				Title:       info.Title,
				Author:      info.Author,
				Content:     info.Content,
				PublishedAt: info.Date,
			}); err != nil {
				fmt.Printf("failed to save record: %v\n", err.Error())
				continue
			}

			fmt.Printf("OK: %v\n", link)
		}
	}
	return nil
}
