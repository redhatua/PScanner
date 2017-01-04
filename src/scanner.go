package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type ScanPortResult struct {
	Port  string
	State bool
	Error error
}

func main() {
	var host = "127.0.0.1"
	var ports = "22,80,443,8080"

	if "" == host || "" == ports {
		log.Fatal("Die!")
	}
	results := []ScanPortResult{}
	portSplit := strings.Split(ports, ",")
	for _, port := range portSplit {
		results = append(results, scanPort(host, port))
	}
	for _, state := range results {
		log.Printf("Port %s %s", state.Port, map[bool]string{true: "open", false: "closed"}[state.State])
	}
}

func scanPort(h, p string) ScanPortResult {
	scan, err := net.Dial("tcp", fmt.Sprintf("%s:%s", h, p))
	res := ScanPortResult{
		Port:  p,
		State: err == nil,
		Error: err,
	}
	if scan != nil {
		scan.Close()
	}
	return res

}
