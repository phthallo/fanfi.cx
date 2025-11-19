package internal

import (
	"github.com/gocolly/colly"
	"fmt"
	"strconv"
)

type Chapter struct {
	ID     int 
	Title  string
	Summary string
	AuthorNotes string
	Content string
}

type Props struct {
	WorkId string
	Chapter int
}

func obtainChapters(work_id string) (chapters []int) {
	var chapter_ids []int
	
	c := colly.NewCollector() // to do: figure out better ? multithreaded colly solution for cases of high use
	//
	c.OnHTML("#selected_id > option", func(e *colly.HTMLElement) {
		fmt.Println("stuff", e)
		chapter_id, err := strconv.Atoi(e.Attr("value"))
		if err != nil {
			fmt.Println("Scraper error, unable to convert chapter ID into number:", err)
		}
		chapter_ids = append(chapter_ids, chapter_id)
	})

	c.Visit(fmt.Sprintf("https://archiveofourown.org/works/%s", work_id))
	return chapter_ids
}

func ScrapeWork(work_id string, chapter int) (*Chapter, error) { // to do: make all numbers actually Numbers
	var chapter_id int
	var work_chapter *Chapter
	var chapter_ids = obtainChapters(work_id)
	c := colly.NewCollector() 

	if len(chapter_ids) != 0 { //multi chapter fic
		if (chapter - 1) > len(chapter_ids) {
			return nil, fmt.Errorf("Invalid chapter number!")
		} else {
			chapter_id = chapter_ids[chapter - 1]
		}
	} else {
		chapter_id = 1
	}

	c.OnHTML("#chapters", func(e *colly.HTMLElement) {
		work_chapter = &Chapter {
			ID:			 chapter_id,
			Title: 		 e.ChildText("div > .chapter.preface.group > h3"),
			Summary:     e.ChildText(".chapter.preface.group > #summary > blockquote"),
			Content:     e.ChildText("div.userstuff"),
			AuthorNotes: e.ChildText(".end.notes.module > blockquote"),
		}
	})

	c.Visit(fmt.Sprintf("https://archiveofourown.org/works/%v/chapters/%v", work_id, chapter_id))
	return work_chapter, nil
}