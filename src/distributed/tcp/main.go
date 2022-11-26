package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type frame struct {
	lines    []string
	duration time.Duration
}

var frames []frame

func Stream(conn net.Conn) {
	var bb bytes.Buffer
	defer conn.Close()

	for _, frame := range frames {
		bb.Reset()
		// Clear terminal and move cursor to position (1,1)
		_, _ = bb.WriteString("\033[2J\033[1;1H")

		for _, line := range frame.lines {
			_, _ = bb.WriteString(line)
			_, _ = bb.WriteString("\r\n")
		}
		_, err := conn.Write(bb.Bytes())
		if err != nil {
			return
		}
		time.Sleep(frame.duration * time.Second / 15)
	}
}

func main() {
	moviePtr := flag.String("movie", "sw1.txt", "path to ASCII movie file")
	addrPtr := flag.String("addr", ":54321", "TCP address to listen on")
	flag.Parse()
	data, err := os.ReadFile(*moviePtr)
	if err != nil {
		fmt.Printf("Failed to load file %s\n", *moviePtr)
	}
	lines := strings.Split(string(data), "\n")
	frameHeight := 13
	for i := range lines {
		if i%(frameHeight+1) == 0 {
			frameDurationStr := lines[i]
			frameDurationInt, err := strconv.ParseInt(frameDurationStr, 0, 64)
			if err != nil {
				fmt.Printf("Failed to parse frame duration from line: %s", frameDurationStr)
			}
			frames = append(frames, frame{lines[i+1 : i+1+frameHeight], time.Duration(frameDurationInt)})
		}
	}
	fmt.Printf("Extracted %d frames from %s\n", len(frames), *moviePtr)

	listen, err := net.Listen("tcp", *addrPtr)
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go Stream(conn)
	}
}
