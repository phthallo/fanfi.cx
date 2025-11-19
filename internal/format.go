package internal

import (
	"fmt"
	"strings"
)

func splitStrings(block string) ([]string){
    var txtStrings []string
    var result string
    lines := strings.Split(block, "\n")
    replacer := strings.NewReplacer(
        "\010", "",
        "\u2019", "'",
        "\u2018", "'",
        "\u2014", "-",  
        "\u2013", "-",
        "\u0222", `'`, // replace all double quotes with single quotes bc ts sucks
        "\u201C", `'`,
        "\u201D", `'`,
        "\u00a0", " ", 
        )
for _, line := range lines {
    for i := 0; i < len(line); {
        end := i + 255
        if end > len(line) {
            end = len(line)
        } else {
            lastDot := strings.LastIndex(line[i:end], ".")
            if lastDot != -1 {
                end = i + lastDot + 1  
            }
        }
        result = replacer.Replace(line[i:end])
        txtStrings = append(txtStrings, result)
        i = end
        }
    }    
    return txtStrings
}

func FormatSearchResults(works []Work) (formatted []string){
    var fullFormattedStrings []string
	for _, work := range works {
		var modifiedDescription []string
		splitLinesDescription := strings.Split(work.Description, "\n")
		for _, line := range splitLinesDescription {
			modifiedDescription = append(modifiedDescription, fmt.Sprintf("|| %s", line))
		}
		block := fmt.Sprintf(`
||=========================
||==================
|| %s by %s
||==================
|| >> Description
%s
||==================
|| >> ID
|| %s
||=========================
.
.`, work.Title, work.Author, strings.Join(modifiedDescription, "\n"), work.ID)
    fullFormattedStrings = append(fullFormattedStrings, splitStrings(block)...)
    }    
    return fullFormattedStrings

}

func FormatWork(chapter *Chapter) (formatted []string){
    block := fmt.Sprintf(`
||=========================
|| %s
||==================
|| >> Summary
|| %s
||=========================
%s
||=========================
|| >> Author Notes
|| %s
||=========================

`, chapter.Title, chapter.Summary, chapter.Content, chapter.AuthorNotes)
    fmt.Println(chapter.Content)
    txtStrings := splitStrings(block)
    return txtStrings
}