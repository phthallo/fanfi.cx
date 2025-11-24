package internal

import (
    "fmt"
    "strings"
    "github.com/phthallo/fanfi.cx/pkg/plaintui"
)

func FormatSearchResults(works []Work) []string {    
    var result []string
    for index, work := range works {
        content := plaintui.Rect([]string{
            fmt.Sprintf("%s by %s", work.Title, work.Author),
            ">> Tags",
            work.Tags,
            ">> Description",
            work.Description, 
            ">> ID",
            work.ID,
        }, 50, 1)
        fmt.Println("Just formatted work", index)
        result = append(result, strings.Split(content, "\n")...)
    }
    fmt.Println("Formatting done! returning")
    return result
}

func FormatWork(chapter *Chapter) []string {
    separate := len(strings.Split(chapter.Content, "\n"))
    fmt.Println("in format work: there are ", separate, "[aragra[hs]]")
    content := plaintui.Rect([]string{
        chapter.ParentTitle + " by " + chapter.Author,
        chapter.Title,
        ">> Summary",
        chapter.Summary, 
        chapter.Content,
        ">> Author Notes",
        chapter.AuthorNotes,
    }, 50, 1)
    
    return strings.Split(content, "\n")
}