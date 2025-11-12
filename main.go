package main

import (
	"github.com/256dpi/newdns"
	"github.com/miekg/dns"
	"fmt"
	"github.com/phthallo/fanfi.cx/internal"
)


func main(){
	zone := &newdns.Zone{
		Name:             "fanfi.cx.",
        MasterNameServer: "ns1.hostmaster.com.",
        AllNameServers: []string{
            "ns1.hostmaster.com.",
            "ns2.hostmaster.com.",
            "ns3.hostmaster.com.",
        },
		Handler: internal.Handler,
	}

server := newdns.NewServer(newdns.Config{
    Handler: func(name string) (*newdns.Zone, error) {
        fmt.Println("Server handler received name:", name)

        if newdns.InZone("fanfi.cx.", name) {
            return zone, nil
        }

        return nil, nil
    },
    Logger: func(e newdns.Event, msg *dns.Msg, err error, reason string) {
        fmt.Println(e, err, reason)
    },
})

go func() {
    err := server.Run(":1337")
    if err != nil {
        panic(err)
    }
}()

fmt.Println("Query apex: dig fanfi.cx @0.0.0.0 -p 1337")
fmt.Println("Query other: dig foo.fanfi.cx @0.0.0.0 -p 1337")

select {}
}