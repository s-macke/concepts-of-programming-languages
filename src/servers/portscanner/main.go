package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	//fmt.Printf("Scan Port %d\n", port)
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)

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
	fmt.Printf("Port Scanning from %d to %d\n", *portMin, *portMax)

	var wg sync.WaitGroup
	for port := *portMin; port <= *portMax; port++ {
		wg.Add(1)
		port := port
		go func() {
			defer wg.Done()
			open := scanPort("tcp", *host, port)
			if open {
				fmt.Printf("Port %d Open\n", port)
			}
		}()
	}
	wg.Wait()
}
