package main

import (
	"testing"
)

func TestLeaseFromText(t *testing.T) {
	// Test if string gets converter correctly
	var text1 = "1623926767 10:10:30:40:50:60 10.10.1.2 some-host 01:10:10:30:40:50:60"
	var text2 = "1623926767 10:10:30:40:50:60 10.10.1.2 some-host 01:10:10:30:40:50:60 some extra junk"
	var text3 = "too few columns"
	var lease1, err1 = leaseFromText(text1)
	var _, err2 = leaseFromText(text2)
	var _, err3 = leaseFromText(text3)

	if err1 != nil {
		t.Errorf("Correct text not pared correctly")
	}

	if lease1.ExpiryTime != 1623926767 {
		t.Errorf("Expiry time not pared correctly, expected: %d, got %d", 1623926767, lease1.ExpiryTime)
	}
	if lease1.MacAddress != "10:10:30:40:50:60" {
		t.Errorf("MAC-Address not pared correctly, expected: %s, got %s", "10:10:30:40:50:60", lease1.MacAddress)
	}
	if lease1.Ip != "10.10.1.2" {
		t.Errorf("IP not pared correctly, expected: %s, got %s", "10.10.1.2", lease1.Ip)
	}
	if lease1.Hostname != "some-host" {
		t.Errorf("Hostname not pared correctly, expected: %s, got %s", "some-host", lease1.Hostname)
	}
	if lease1.ClientID != "01:10:10:30:40:50:60" {
		t.Errorf("Client ID not pared correctly, expected: %s, got %s", "01:10:10:30:40:50:60", lease1.ClientID)
	}

	if err2 == nil {
		t.Errorf("Text with extra columns parsed, expected nil")
	}

	if err3 == nil {
		t.Errorf("Text with too few columns parsed, expected nil")
	}
}

func TestFileParser(t *testing.T) {
	//
	var fileUrl = "test/dnsmasq.leases"
	var leases, err = parseLeasesFile(&fileUrl)

	if err != nil {
		t.Errorf("Could not read test file.")
	}

	if len(leases) != 2 {
		t.Errorf("number of parsed leases is %d, expected 2", len(leases))
	}

	if leases[0].ExpiryTime != 1623926766 {
		t.Errorf("Expiry time of first test record not pared correctly, expected: %d, got %d", 1623926766, leases[0].ExpiryTime)
	}

	if leases[1].Hostname != "some-host" {
		t.Errorf("Hostname of first test record not pared correctly, expected: %s, got %s", "some-host", leases[1].Hostname)
	}
}

func TestFileParserWithNonExistingFile(t *testing.T) {
	//
	var fileUrl = "test/no/such/file"
	var _, err = parseLeasesFile(&fileUrl)
	if err == nil {
		t.Errorf("Reading non existent file should return error")
	}
}

func TestEncodeLeasesToJson(t *testing.T) {
	leases := make([]Lease, 2)
	_, err := encodeLeasesToJson(leases)
	if err != nil {
		t.Errorf("Could not encode slice of leases to json")
	}
}
