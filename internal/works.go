package internal

import (
	"github.com/gocolly/colly"
	"fmt"
	"strconv"
	"context"
	"time"
	"strings"
)

type Chapter struct {
	ID     int 
	ParentTitle string
	Author string
	Title  string
	Summary string
	AuthorNotes string
	Content string
}

type Props struct {
	WorkId string
	Chapter int
}

func obtainChapter(work_id string, chapter int) (ch int) {	
	var chapter_id int
	var err error
	c := colly.NewCollector() // to do: figure out better ? multithreaded colly solution for cases of high use
    c.SetRequestTimeout(5 * time.Second)
    c.AllowURLRevisit = true

	//
	fmt.Println("Chapter was ", fmt.Sprintf("#selected_id > option:nth-child(%v)", chapter))

	c.OnHTML(fmt.Sprintf("#selected_id > option:nth-child(%v)", chapter), func(e *colly.HTMLElement) {
		fmt.Println("stuff", e)
		chapter_id, err = strconv.Atoi(e.Attr("value"))
		if err != nil {
			fmt.Println("Scraper error, unable to convert chapter ID into number:", err)
		}
	})

	c.Visit(fmt.Sprintf("https://archiveofourown.org/works/%s", work_id))
	return chapter_id
}

func ScrapeWork(work_id string, chapter int) (*Chapter, error) { // to do: make all numbers actually Numbers
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	result := make(chan *Chapter, 1)

	var chapter_id int
	var work_chapter *Chapter
	fmt.Println("Chapter passed to scrape work was", chapter)
	if chapter != 1 {
		chapter_id = obtainChapter(work_id, chapter)
		fmt.Println("Chapter ID obtained", chapter_id)
	} else {
		chapter_id = 0
	}
	c := colly.NewCollector()
    c.SetRequestTimeout(5 * time.Second)
    c.AllowURLRevisit = true



	//if len(chapter_ids) != 0 { //multi chapter fic
	//	if (chapter - 1) > len(chapter_ids) {
	//		return nil, fmt.Errorf("Invalid chapter number!")
	//	} else {
	//		chapter_id = chapter_ids[chapter - 1]
	//	}
	//} 
	    
    go func() {
		fmt.Println("Goroutine running")
		c.OnHTML("#workskin", func(e *colly.HTMLElement) {
			var paragraphs []string
			e.ForEach("#chapters > div", func(_ int, el *colly.HTMLElement) {
				paragraphs = append(paragraphs, el.Text)
			})
			text := strings.Join(paragraphs, "\n")

			work_chapter = &Chapter {
				ID:			 chapter_id,
				ParentTitle: e.ChildText("div.preface.group > h2.title.heading"),
				Author:      e.ChildText("div.preface.group > h3.byline.heading"),
				Title: 		 e.ChildText("#chapters > div > .chapter.preface.group > h3"),
				Summary:     e.ChildText("#chapters > chapter.preface.group > #summary > blockquote"),
				Content:     text,
				AuthorNotes: e.ChildText("#chapters > chapter.preface.group"),
			}
		})

		workToScrape := fmt.Sprintf("https://archiveofourown.org/works/%v", work_id)
		if chapter_id != 0 {
			workToScrape += fmt.Sprintf("/chapters/%v", chapter_id)
		}
		fmt.Println("Visiting ", workToScrape)
		err := c.Visit(workToScrape)
		fmt.Println("scraping work error", err)
		result <- work_chapter
	}()


    select {
		case work := <-result:
			return work, nil
		case <-ctx.Done():
			return nil, fmt.Errorf("scrape timeout")
    }

}