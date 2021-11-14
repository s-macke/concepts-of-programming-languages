package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	host := flag.String("host", "localhost", "hostname")
	portMin := flag.Int("portmin", 80, "port min")
	portMax := flag.Int("portmax", 80, "port max")
	flag.Parse()
	fmt.Printf("Port Scanning of host %s from %d to %d\n", *host, *portMin, *portMax)

	for port := *portMin; port <= *portMax; port++ {
		address := *host + ":" + strconv.Itoa(port)
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err == nil {
			fmt.Printf("Port %d Open\n", port)
			conn.Close()
		}
	}
}
