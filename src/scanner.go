package main

import (
	"fmt"
	"net"
	"strings"
)

type ScanPortResult struct {
	Port  int
	State bool
	Error error
}

type Result struct {
	Results []ScanPortResult
}

func main() {
	var host = "127.0.0.1"
	var ports = "22,80,443,8080"

	if "" != host && "" != ports {
		portSplit := strings.Split(ports, ",")
		for port := range portSplit {
			//scan and add to slice
		}
		//make results
	}
}

func appendResult() {
	/* TODO: investigate how we can append it*/
}

func scanPort(h string, p int) *ScanPortResult {
	scan, err := net.Dial("tcp", fmt.Sprintf("%s:%d", h, p))
	res := ScanPortResult{
		Port:  p,
		State: err == nil,
		Error: err,
	}
	if scan != nil {
		scan.Close()
	}
	return &res

}
