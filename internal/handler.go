package internal 

import (
	"github.com/256dpi/newdns"
)

func Handler(name string) ([]newdns.Set, error) {
			if name == ""  {
				return []newdns.Set{
					{
						Name: "fanfi.cx.",
						Type: newdns.A,
						Records: []newdns.Record{
							{Address: "1.2.3.4"},
						},
					},
					{
						Name: "fanfi.cx.",
						Type: newdns.AAAA,
						Records: []newdns.Record{
							{Address: "1:2:3:4::"},
						},
					},
				}, nil
			}

			if name == "foo" {
				return []newdns.Set{
					{
						Name: "foo.fanfi.cx.",
						Type: newdns.CNAME,
						Records: []newdns.Record{
							{Address: "bar.fanfi.cx."},
						},
					},
				}, nil
			}

			return nil, nil
		}