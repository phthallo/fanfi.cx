package internal

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
	"time"
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
	
	var works []Work
	var found bool
	var lengthToTruncateTo int 
	var lengthTags int

	c := colly.NewCollector()
    c.SetRequestTimeout(10 * time.Second)
    c.AllowURLRevisit = true

	c.OnHTML("li[class]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("class"), "work blurb group") {
			// we've found a work!
			var author string
			var tags string
			author = e.ChildText("div.header > h4.heading > a:nth-child(2)")
			if author == " " { 
				author = "Anonymous"
			}
			tags = e.ChildText("ul.tags.commas")
			lengthTags = len(tags) 
			if lengthTags > 150 {
				lengthToTruncateTo = 150
			} else {
				lengthToTruncateTo = lengthTags
			}
			fmt.Println("Debug: work was ", e.ChildText("div.header > h4.heading > a:nth-child(1)"))
			work := Work{
				ID:          strings.Split(e.Attr("id"), "_")[1],
				Title:       e.ChildText("div.header > h4.heading > a:nth-child(1)"),
				Author:      author,
				Tags:		 tags[:lengthToTruncateTo] + "...",
				Description: e.ChildText("blockquote"),
			}
			works = append(works, work)
			found = true
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(fmt.Sprintf("https://archiveofourown.org/works/search?work_search%%5Bquery%%5D=%s&work_search%%5Btitle%%5D=&work_search%%5Bcreators%%5D=&work_search%%5Brevised_at%%5D=&work_search%%5Bcomplete%%5D=&work_search%%5Bcrossover%%5D=&work_search%%5Bsingle_chapter%%5D=0&work_search%%5Bword_count%%5D=&work_search%%5Blanguage_id%%5D=en&work_search%%5Bfandom_names%%5D=&work_search%%5Brating_ids%%5D=&work_search%%5Bcharacter_names%%5D=&work_search%%5Brelationship_names%%5D=&work_search%%5Bfreeform_names%%5D=&work_search%%5Bhits%%5D=&work_search%%5Bkudos_count%%5D=&work_search%%5Bcomments_count%%5D=&work_search%%5Bbookmarks_count%%5D=&work_search%%5Bsort_column%%5D=_score&work_search%%5Bsort_direction%%5D=desc&commit=Search", query))

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("No works found!")
	}
	return works, nil
}