package plaintui 
// barebones tui framework

import (
	"strings"
	"fmt"
	"math"
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

func splitStrings(block string, width int, padding int) []string {
    lines := strings.Split(block, "\n")
    txtStrings := make([]string, 0, len(lines))
    
    for _, line := range lines {
        for len(line) > 0 {
            end := width - 4 - padding
            if end > len(line) {
                end = len(line)
            } else {
				locationOfLastFullStop := float64(strings.LastIndex(line[:end], "."))
				locationOfLastSpace := float64(strings.LastIndex(line[:end], " "))
				idx := math.Max(locationOfLastFullStop, locationOfLastSpace)
                end = int(idx) + 1
            } 
				
            filtered := replacer.Replace(line[:end]) 

			whitespaceNeeded := (width - 4) - len(filtered)

			filtered = strings.Repeat(" ", padding) + filtered + strings.Repeat(" ", whitespaceNeeded)
            txtStrings = append(txtStrings, filtered)
            line = line[end:]
        }
    }
    return txtStrings
}


func Rect(content []string, width int, padding int) (formatted string){
    var buf strings.Builder
	buf.WriteString(fmt.Sprintf("\n||%v||\n", strings.Repeat("=", width - 4 + padding))) // top 
	var verticalLinePadding int
	if padding % 2 == 0 {
		verticalLinePadding = padding / 2
	} else {
		verticalLinePadding = (padding - 1)/2
	}
	for index, partialContent := range content {
		for _ = range verticalLinePadding {
				buf.WriteString(fmt.Sprintf("||%v||\n", strings.Repeat(" ", width - 4 + padding))) // padding for start of section
		}
		splitContent := splitStrings(partialContent, width, padding)
		
		for _, line := range splitContent {
			buf.WriteString(fmt.Sprintf("||%v||\n", line)) // actual content
		}

		for _ = range verticalLinePadding {
			buf.WriteString(fmt.Sprintf("||%v||\n", strings.Repeat(" ", width - 4 + padding))) // padding for end of section
		}

		if index != len(content) - 1{
			buf.WriteString(fmt.Sprintf("||%v||\n", strings.Repeat("=", width - 4 + padding))) // divider
		}
	}

	buf.WriteString(fmt.Sprintf("||%v||\n", strings.Repeat("=", width - 4 + padding))) // bottom
	return buf.String()
}

