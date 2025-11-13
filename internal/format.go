package internal

import (
	"fmt"
	"strings"
)


func FormatSearchResults(works []Work) (formatted []string){
	var txtStrings []string
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
        lines := strings.Split(block, "\n")
        for _, line := range lines {
            for i := 0; i < len(line); i += 255 {
                end := i + 255
                if end > len(line) {
                    end = len(line)
                }
                txtStrings = append(txtStrings, line[i:end])
            }
        }
    }
    return txtStrings

}