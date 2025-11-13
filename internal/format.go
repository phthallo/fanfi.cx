package internal

import (
	"fmt"
	"strings"
)


func FormatSearchResults(works []Work) (formatted []string){
	var txtStrings []string
	for _, work := range works {
		block := fmt.Sprintf(`
|--------------------
||----------------
|| %s 
||----------------
|| >> Description
%s
||----------------
|| >> ID
|| %s
|--------------------
.
.`, work.Title, work.Description, work.ID)
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