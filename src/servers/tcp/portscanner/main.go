package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanPort(hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	host := flag.String("host", "localhost", "hostname")
	portMin := flag.Int("portmin", 80, "port min")
	portMax := flag.Int("portmax", 80, "port max")
	flag.Parse()
	fmt.Printf("Port Scanning of host %s from %d to %d\n", *host, *portMin, *portMax)
	var wg sync.WaitGroup
	for port := *portMin; port <= *portMax; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			if scanPort(*host, port) {
				fmt.Printf("Port %d Open\n", port)
			}
		}(port)
	}
	wg.Wait()
}
