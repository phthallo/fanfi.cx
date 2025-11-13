package internal 

import (
	"github.com/256dpi/newdns"
	"fmt"
	"regexp"
)

func sanitiseDNSLabel(label string) string {
    re := regexp.MustCompile(`[^a-zA-Z0-9-]`)
    return re.ReplaceAllString(label, "-")
}

func Handler(name string) ([]newdns.Set, error) {
		var query []Work
		var err error
		var fqdn = name + "fanfi.cx."
		var newName = sanitiseDNSLabel(name)

		if newName != "" {
			query, err = QuerySearchResults(newName)
			if err != nil {
				fmt.Println("this is where the err", err)
			}

			return []newdns.Set{
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
				{ 
					Name: fqdn,
					Type: newdns.TXT,
					Records: []newdns.Record{
						{Data: []string{query[0].Title}},
					},
				},
			}, nil
		}

			return nil, nil
		}