package main

import (
	"flag"
	"fmt"
)

var (
	listen = flag.String("listen",
		"0.0.0.0:9154",
		"Address to listen on")

	leasesPath = flag.String("leases_path",
		"/var/lib/misc/dnsmasq.leases",
		"Path to dnsmasq leases file")
)

type Lease struct {
	expiryTime int
	macAddress string
	ip         string
	hostname   string
	clientID   string
}

func parseLeasesFile(fileUrl *string) []Lease {
	var leases []Lease
	return leases
}

func main() {
	flag.Parse()

	fmt.Println("Listening on ", *listen)
}
