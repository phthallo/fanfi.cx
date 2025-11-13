package internal 

import (
	"github.com/256dpi/newdns"
	"fmt"
	"regexp"
)

func sanitiseDNSLabel(label string) string {
	label = regexp.MustCompile(`[\\ ]`).ReplaceAllString(label, "+")
	re := regexp.MustCompile(`[^a-zA-Z0-9+._~-]`)
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

			txtSet := newdns.Set{
				Name: fqdn,
				Type: newdns.TXT,
				Records: []newdns.Record{},
			}
			for _, txt := range FormatSearchResults(query) {
				txtSet.Records = append(txtSet.Records, newdns.Record{Data: []string{txt}})
			}
			sets = append(sets, txtSet)
			return sets, nil
		}

			return nil, nil
		}