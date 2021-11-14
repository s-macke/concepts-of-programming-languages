package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanPort(hostname string, wg *sync.WaitGroup, portCh chan int) {
	defer wg.Done()
	for {
		port, ok := <-portCh
		if !ok {
			break
		}
		address := hostname + ":" + strconv.Itoa(port)
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err == nil {
			fmt.Printf("Port %d Open\n", port)
			conn.Close()
		}
	}
}

func main() {
	host := flag.String("host", "localhost", "hostname")
	portMin := flag.Int("portmin", 80, "port min")
	portMax := flag.Int("portmax", 80, "port max")
	flag.Parse()
	fmt.Printf("Port Scanning of host %s from %d to %d\n", *host, *portMin, *portMax)

	var wg sync.WaitGroup
	portCh := make(chan int)
	for i := 0; i < 10; i++ { // limit
		wg.Add(1)
		go scanPort(*host, &wg, portCh)
	}

	for port := *portMin; port <= *portMax; port++ {
		portCh <- port
	}

	close(portCh)
	wg.Wait()
}
