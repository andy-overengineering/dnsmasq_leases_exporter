package main

import (
	"testing"
)

func TestFileParser(t *testing.T) {
	//
	var fileUrl = "test/dnsmasq.leases"
	var leases = parseLeasesFile(&fileUrl)
	if len(leases) != 2 {
		t.Errorf("number of parsed leases is %d, expected 2", len(leases))
	}
}
