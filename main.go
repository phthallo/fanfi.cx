package main

import (
	"github.com/256dpi/newdns"
	"github.com/miekg/dns"
	"fmt"
	"github.com/phthallo/fanfi.cx/internal"
    "github.com/joho/dotenv"    
    "os"
)


func main(){
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("Error loading environment variables!")
	}

	zone := &newdns.Zone{
		Name:             os.Getenv("FQDN"),
        MasterNameServer: "ns1.hostmaster.com.",
        AllNameServers: []string{
            "ns1.hostmaster.com.",
            "ns2.hostmaster.com.",
            "ns3.hostmaster.com.",
        },
        AdminEmail: "generic@email.com",
		Handler: internal.Handler,
	}

server := newdns.NewServer(newdns.Config{
    Handler: func(name string) (*newdns.Zone, error) {
        fmt.Println("Server handler received name:", name)
        return zone, nil
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

fmt.Println(`Search AO3: dig @0.0.0.0 "your search here" -p 1337`)

select {}
}