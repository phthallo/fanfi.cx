package main

import (
	"github.com/256dpi/newdns"
	"github.com/miekg/dns"
	"fmt"
	"github.com/phthallo/fanfi.cx/internal"
    "github.com/joho/godotenv"    
    "os"
)


func main(){
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("Error loading environment variables!")
	}

	zone := &newdns.Zone{
		Name:             os.Getenv("FQDN"),
        MasterNameServer: os.Getenv("MASTER_NS"),
        AllNameServers: []string{
            os.Getenv("SECONDARY_NS"),
            os.Getenv("TERTIARY_NS"),
            os.Getenv("QUARTERNARY_NS"),
        },
        AdminEmail: "generic@email.com",
		Handler: internal.Handler,
	}

server := newdns.NewServer(newdns.Config{
    Addr: "0.0.0.0:1337",
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