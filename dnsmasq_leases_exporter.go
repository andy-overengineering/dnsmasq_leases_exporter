package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	listenArg = flag.String("listen",
		"0.0.0.0:9154",
		"Address to listen on")

	leasesPathArg = flag.String("leases_path",
		"/var/lib/misc/dnsmasq.leases",
		"Path to dnsmasq leases file")
)

type Lease struct {
	ExpiryTime int    `json:"expiryTime"`
	MacAddress string `json:"macAddress"`
	Ip         string `json:"ip"`
	Hostname   string `json:"hostname"`
	ClientID   string `json:"clientID"`
}

func leaseFromText(text string) (*Lease, error) {
	// Parse line of text in leases file format and return Lease struct
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	t := ""
	l := Lease{}
	column := 0

	for scanner.Scan() {
		t = scanner.Text()

		// Count column
		column += 1
		switch {
		case column == 1:
			i, err := strconv.Atoi(t)
			if err != nil {
				return &l, errors.New(fmt.Sprintf("Could not parse %s as expiry time\n", t))
			}
			l.ExpiryTime = i
		case column == 2:
			l.MacAddress = t
		case column == 3:
			l.Ip = t
		case column == 4:
			l.Hostname = t
		case column == 5:
			l.ClientID = t

		}
	}

	if column != 5 {

		return &l, errors.New(fmt.Sprintf("Unexpected number of columns in leases file, expected 5, got %d\n", column))
	}

	return &l, nil
}

func parseLeasesFile(fileUrl *string) ([]Lease, error) {
	// Parse given file line by line and return slice of Lease structs for each line that could be parsed
	var leases []Lease
	f, err := os.Open(*fileUrl)

	if err != nil {
		// File could not be read; prit error to output and return a nil

		return leases, errors.New(fmt.Sprintf("Could not read file %s: %s\n", *fileUrl, err))
	}

	// Close file at the end
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l, err := leaseFromText(scanner.Text())
		if err == nil {
			leases = append(leases, *l)
		}
	}

	return leases, nil
}

func encodeLeasesToJson(leases []Lease) (string, error) {
	// build json string from slice of Lease structs
	b, err := json.Marshal(leases)
	if err == nil {
		return string(b), nil
	}
	return "", err
}

type server struct {
	listen     string
	leasesPath string
}

func (s *server) leasesHandler(w http.ResponseWriter, r *http.Request) {
	leases, err := parseLeasesFile(&s.leasesPath)
	if err != nil {
		fmt.Printf("Error delivering leases:\n%s\n\n", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	resp, err := encodeLeasesToJson(leases)

	if err != nil {
		fmt.Printf("Error delivering leases:\n%s\n\n", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, resp)
}

func main() {
	flag.Parse()

	s := &server{
		listen:     *listenArg,
		leasesPath: *leasesPathArg,
	}

	http.HandleFunc("/leases", s.leasesHandler)

	log.Println("Listening on ", s.listen)
	log.Println("Serving leases from ", s.leasesPath)
	log.Fatal(http.ListenAndServe(s.listen, nil))
}
