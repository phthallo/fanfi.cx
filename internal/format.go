package internal

import (
    "fmt"
    "strings"
    "github.com/phthallo/fanfi.cx/pkg/plaintui"
)

var replacer = strings.NewReplacer(
    "\010", "",
    "\u2019", "'",
    "\u2018", "'",
    "\u2014", "-",
    "\u2013", "-",
    "\u0222", "'",
    "\u201C", "'",
    "\u201D", "'",
    "\u00a0", " ",
    "\u2025", "...",
)


func FormatSearchResults(works []Work) []string {    
    var result []string
    for _, work := range works {
        content := plaintui.Rect([]string{
            fmt.Sprintf("%s by %s", work.Title, work.Author),
            ">> Description",
            work.Description, 
            ">> ID",
            work.ID,
        }, 80, 1)
        result = append(result, strings.Split(content, "\n")...)
        
    }
    
    return result
}

func FormatWork(chapter *Chapter) []string {
    fmt.Println(chapter.Content)
    separate := len(strings.Split(chapter.Content, "\n"))
    fmt.Println("in format work: there are ", separate, "[aragra[hs]]")
    content := plaintui.Rect([]string{
        chapter.Title,
        ">> Summary",
        chapter.Summary, 
        chapter.Content,
        ">> Author Notes",
        chapter.AuthorNotes,
    }, 80, 1)
    
    return strings.Split(content, "\n")
}