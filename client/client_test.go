package main

import (
	"log"
	"testing"
)

func TestAllAvailbleIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		requestOption = new(int)
		*requestOption = 1
		var respMsg, respStatus = requestSelection(requestOption)

		log.Printf("TESTING AllAvailbleIPs: Returned Resp Status code: %v. Returned Resp length: %v", respStatus, len(respMsg))

		if respStatus != "200 OK" {
			t.Errorf("Returned status is not 200 OK. Got %v", respStatus)
		}

		if len(respMsg) < 2 {
			t.Errorf("Returned JSON is empty (Char less than 2)")
		}
	}
}
func TestGetIP(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		requestOption = new(int)
		*requestOption = 2
		var respMsg, respStatus = requestSelection(requestOption)

		log.Printf("TESTING GetIP: Returned Resp Status code: %v. Returned Resp length: %v", respStatus, len(respMsg))

		if respStatus != "200 OK" {
			t.Errorf("Returned status is not 200 OK. Got %v", respStatus)
		}

		if len(respMsg) < 2 {
			t.Errorf("Returned JSON is empty (Char less than 2)")
		}
	}
}

func TestDeleteIPfromPool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		requestOption = new(int)
		*requestOption = 3
		var respMsg, respStatus = requestSelection(requestOption)

		log.Printf("TESTING DeleteIPfromPool: Returned Resp Status code: %v. Returned Resp length: %v", respStatus, len(respMsg))

		if respStatus != "200 OK" {
			t.Errorf("Returned status is not 200 OK. Got %v", respStatus)
		}

		if respMsg != "a-102.131.46.22 IP deleted " {
			t.Errorf("Returned response message is incorrect. Got : %v", respMsg)
		}
	}
}

func TestAddIPtoPool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		requestOption = new(int)
		*requestOption = 4
		var respMsg, respStatus = requestSelection(requestOption)

		log.Printf("TESTING AddIPtoPool: Returned Resp Status code: %v. Returned Resp length: %v", respStatus, len(respMsg))

		if respStatus != "200 OK" {
			t.Errorf("Returned status is not 200 OK. Got %v", respStatus)
		}

		if respMsg != "New IP posted" {
			t.Errorf("Returned response message is incorrect. Got : %v", respMsg)
		}
	}
}

func TestCreateNewIPpool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
		requestOption = new(int)
		*requestOption = 5
		var respMsg, respStatus = requestSelection(requestOption)

		log.Printf("TESTING CreateNewIPpool: Returned Resp Status code: %v. Returned Resp length: %v", respStatus, len(respMsg))

		if respStatus != "200 OK" {
			t.Errorf("Returned status is not 200 OK. Got %v", respStatus)
		}

		if respMsg != "IP address a-253.14.93.192 changed to a-111.11.11.111" {
			t.Errorf("Returned response message is incorrect. Got : %v", respMsg)
		}
	}
}
