package internal 

import (
	"github.com/256dpi/newdns"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

type Search struct {
	WorkId string
	SearchQuery string
	Chapter int 
}

func validateAndSanitiseDNSLabel(label string) (map[string]string, error) {
	label = strings.ReplaceAll(label, `\ `, ` `) // dns thingy will always escape characters
	params := regexp.MustCompile(`\[(\w+)\]\s+([^\[]+)`).FindAllStringSubmatch(label, -1)	
	fmt.Println(params)	

	validParams := make(map[string]string)

	for i, match := range params {
		tag := match[1]
		phrase := strings.TrimSpace(match[2])
		
		if (tag == "chapter"){
			if i == 0 { 
				continue 
			}
			prevTag := params[i-1][1] 
			if (prevTag != "work_id" && prevTag != "search") {
				continue 
			}

			if _, err := strconv.Atoi(phrase); err != nil {		
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
	fmt.Println(validParams)
	return validParams, nil
}

func Handler(name string) ([]newdns.Set, error) {
		var searchResults []Work
		var fqdn = name + "fanfi.cx."
		var parsedSearchParams, err = validateAndSanitiseDNSLabel(name)
		if err != nil {
			fmt.Println(err)
		}
		var searchQuery = parsedSearchParams["search"]
		fmt.Println("Search query",  parsedSearchParams["search"], "Chapter",  parsedSearchParams["chapter"], "WorkId",  parsedSearchParams["work_id"])
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

		if searchQuery != "" {
			searchResults, err = QuerySearchResults(searchQuery)
			if err != nil {
				fmt.Println("this is where the err", err)
			}

			for _, txt := range FormatSearchResults(searchResults) {
				txtSet.Records = append(txtSet.Records, newdns.Record{Data: []string{txt}})
			}
		}
		sets = append(sets, txtSet)
		return sets, nil
		}