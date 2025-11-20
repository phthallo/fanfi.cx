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
		Name:             ".",
        MasterNameServer: os.Getenv("PRIMARY_NS"),
        AllNameServers: []string{
            os.Getenv("PRIMARY_NS"),
            os.Getenv("SECONDARY_NS"),
            os.Getenv("TERTIARY_NS"),
            os.Getenv("QUATERNARY_NS"),
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
    port := os.Getenv("PORT")

    if port == "" {
        port = "53"
    }
    fmt.Println(`Server is up and running on port`, port)

    err := server.Run(fmt.Sprintf(":%v", port))
    if err != nil {
        panic(err)
    }
}()


select {}
}