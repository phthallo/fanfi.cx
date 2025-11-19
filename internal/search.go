package internal

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
)

type Work struct {
	ID string
	Title string
	Author string
	Description string
	Fandom string
	Date string
	Tags string
	Language string
	Words string
	ChaptersCount string
	Comments string
	Kudos string
	Bookmarks string
	Hits string
}

func QuerySearchResults(query string) ([]Work, error) {
	c := colly.NewCollector()
	var works []Work
	var found bool

	c.OnHTML("li[class]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("class"), "work blurb group") {
			// we've found a work!
			var author string
			author = e.ChildText("div.header > h4.heading > a:nth-child(2)")
			if author == " " { 
				author = "Anonymous"
			}
			work := Work{
				ID:          strings.Split(e.Attr("id"), "_")[1],
				Title:       e.ChildText("div.header > h4.heading > a:nth-child(1)"),
				Author:      author,
				Description: e.ChildText("blockquote"),
			}
			works = append(works, work)
			found = true
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(fmt.Sprintf("https://archiveofourown.org/works/search?=work_search[language_id]=en&work_search%%5Bquery%%5D=%s", query))
	
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("No works found!")
	}
	return works, nil
}