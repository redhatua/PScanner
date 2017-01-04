package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type ScanPortResult struct {
	Port  string
	State string
	Error error
}

const (
	PortOpen     = "Open"
	PortClosed   = "Closed"
	PortFiltered = "Filtered"
)

func main() {
	var host = flag.String("h", "", "Target host")
	var ports = flag.String("p", "", "Port(s) for scan")
	flag.Parse()
	if "" == *host || "" == *ports {
		log.Fatal("Die!")
	}
	results := []ScanPortResult{}
	portSplit := parsePorts(*ports)
	for _, port := range portSplit {
		results = append(results, scanPort(*host, port))
	}
	showResult(results)
}

func parsePorts(ports string) []string {
	var result = []string{}
	split := strings.Split(ports, ",")
	for _, port := range split {
		if strings.Contains(port, "-") {
			if r, err := splitRange(port); err == nil {
				result = append(result, r...)
			}
		} else {
			result = append(result, port)
		}
	}
	return result
}

func splitRange(r string) ([]string, error) {
	result := []string{}
	split := strings.SplitN(r, "-", 2)
	min, _ := strconv.Atoi(split[0])
	max, _ := strconv.Atoi(split[1])
	for i := min; i <= max; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result, nil
}

func scanPort(h, p string) ScanPortResult {
	scan, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", h, p), time.Second)
	state := PortOpen
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "timeout"):
		state = PortFiltered
	case strings.Contains(err.Error(), "refused"):
		state = PortClosed
	}
	res := ScanPortResult{
		Port:  p,
		State: state,
		Error: err,
	}
	if scan != nil {
		scan.Close()
	}
	return res
}

func showResult(results []ScanPortResult) {
	log.Printf("Scanned %d ports", len(results))
	for _, state := range results {
		if state.State != PortClosed {
			log.Printf("Port %s %s", state.Port, state.State)
		}
	}
}
