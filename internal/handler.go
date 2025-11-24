package internal 

import (
	"github.com/256dpi/newdns"
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"os"
)

type Search struct {
	WorkId string
	SearchQuery string
	Chapter int 
}

func validateAndSanitiseDNSLabel(label string) (map[string]string, error) {
	label = strings.ReplaceAll(label, `\ `, ` `) // dns thingy will always escape characters
	params := regexp.MustCompile(`\[(\w+)\]\s+([^\[]+)`).FindAllStringSubmatch(label, -1)	

	validParams := make(map[string]string)


	var tag string
	var phrase string
	var err error

	if (params == nil){
		phrase = regexp.MustCompile(`[\\ ]`).ReplaceAllString(label, "+") 
		phrase = regexp.MustCompile(`[^a-zA-Z0-9+._~-]`).ReplaceAllString(label, "+")
		validParams["search"] = phrase
	}

	for i, match := range params {
		tag = match[1]
		phrase = strings.TrimSpace(match[2])
		
		if (tag == "chapter"){
			if i == 0 { 
				continue 
			}
			prevTag := params[i-1][1] 
			if (prevTag != "work_id" && prevTag != "search") {
				continue 
			}

			if _, err = strconv.Atoi(phrase); err != nil {		
				return nil, fmt.Errorf("Provided chapter is not a number!")
			}
		}
		
		if tag == "search" {
			if (i+1 < len(params) && params[i+1][1] == "chapter") {
				continue 
			} else {
				// make url friendly
				phrase = regexp.MustCompile(`[\\ ]`).ReplaceAllString(phrase, "+") 
				phrase = regexp.MustCompile(`[^a-zA-Z0-9+._~-]`).ReplaceAllString(phrase, "-")
			}
		}
		
    validParams[tag] = phrase
	}

	return validParams, nil
}

func Handler(name string) ([]newdns.Set, error) {
		var searchResults []Work
		var chapterResults *Chapter
		var parsedSearchParams, err = validateAndSanitiseDNSLabel(name)
		var fqdn string
		fqdn = regexp.MustCompile(`[^a-zA-Z0-9._~-]+`).ReplaceAllString(fmt.Sprint(parsedSearchParams), ".")
		if os.Getenv("FQDN") != "." {
 			fqdn += os.Getenv("FQDN")
		}
		if err != nil {
			fmt.Println("Handler error", err)
		}
		var searchQuery = parsedSearchParams["search"]

		var chapter int
		if parsedSearchParams["chapter"] != "" && parsedSearchParams["search"] == "" {
			if int_chapter, err := strconv.Atoi(parsedSearchParams["chapter"]); err != nil {
				fmt.Println("Error converting chapter to number", err)
			} else {
				chapter = int_chapter
			}
		} else {
			chapter = 1
		}
		fmt.Println("Determined chapter as", chapter)
		
		var workID = parsedSearchParams["work_id"]
		// var output string

		sets := []newdns.Set{
			{
				Name: fqdn,
				Type: newdns.A,
				Records: []newdns.Record{
					{Address: "1.2.3.4"},
				},
			},
			{
				Name: fqdn,
				Type: newdns.AAAA,
				Records: []newdns.Record{
					{Address: "1:2:3:4::"},
				},
			},
		}

		txtSet := newdns.Set {
				Name: fqdn,
				Type: newdns.TXT,
				Records: []newdns.Record{},
		}

		if searchQuery != "" { // TO DO: add support for defaulting to search query if no tags passed
			searchResults, err = QuerySearchResults(searchQuery)
			if err != nil {
				fmt.Println("this is where the err", err)
			}

			formatted := FormatSearchResults(searchResults)
			fmt.Println(formatted)
			for _, txt := range formatted {
				txtSet.Records = append(txtSet.Records, newdns.Record{Data: []string{txt}})
			}
		}
		
		if workID != "" {
			chapterResults, err = ScrapeWork(workID, chapter)

			if err != nil {
				fmt.Println("Error fetching work:", err)
			} 

			if chapterResults != nil {
				for _, txt := range FormatWork(chapterResults) {
					txtSet.Records = append(txtSet.Records, newdns.Record{Data: []string{txt}})
				}
			}
		}

		sets = append(sets, txtSet)
		return sets, nil
		}