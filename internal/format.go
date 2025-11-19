package internal

import (
    "fmt"
    "strings"
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
)

func splitStrings(block string) []string {
    lines := strings.Split(block, "\n")
    txtStrings := make([]string, 0, len(lines))
    
    for _, line := range lines {
        for len(line) > 0 {
            end := 255
            if end > len(line) {
                end = len(line)
            } else if idx := strings.LastIndex(line[:end], "."); idx != -1 {
                end = idx + 1
            }
            
            txtStrings = append(txtStrings, replacer.Replace(line[:end]))
            line = line[end:]
        }
    }
    return txtStrings
}

func FormatSearchResults(works []Work) []string {
    var buf strings.Builder
    result := make([]string, 0, len(works)*50)
    
    for _, work := range works {
        buf.WriteString("\n||=========================\n")
        buf.WriteString("||==================\n")
        buf.WriteString(fmt.Sprintf("|| %s by %s\n", work.Title, work.Author))
        buf.WriteString("||==================\n")
        buf.WriteString("|| >> Description\n")
        
        for _, line := range strings.Split(work.Description, "\n") {
            buf.WriteString(fmt.Sprintf("|| %s\n", line))
        }
        
        buf.WriteString("||==================\n")
        buf.WriteString(fmt.Sprintf("|| >> ID\n|| %s\n", work.ID))
        buf.WriteString("||=========================\n.\n.\n")
    }
    
    result = append(result, splitStrings(buf.String())...)
    return result
}

func FormatWork(chapter *Chapter) []string {
    var buf strings.Builder
    buf.WriteString("\n||=========================\n")
    buf.WriteString(fmt.Sprintf("|| %s\n", chapter.Title))
    buf.WriteString("||==================\n")
    buf.WriteString(fmt.Sprintf("|| >> Summary\n|| %s\n", chapter.Summary))
    buf.WriteString("||=========================\n")
    buf.WriteString(chapter.Content)
    buf.WriteString("\n||=========================\n")
    buf.WriteString(fmt.Sprintf("|| >> Author Notes\n|| %s\n", chapter.AuthorNotes))
    buf.WriteString("||=========================\n")
    
    return splitStrings(buf.String())
}