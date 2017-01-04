package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type ScanPortResult struct {
	Port  string
	State bool
	Error error
}

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
	for _, state := range results {
		log.Printf("Port %s %s", state.Port, map[bool]string{true: "open", false: "closed"}[state.State])
	}
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
