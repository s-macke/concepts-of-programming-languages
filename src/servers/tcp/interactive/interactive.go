package main

import (
	"bufio"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "simulationcorner.net:54322")
	if err != nil {
		panic(err)
	}
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _, err := reader.ReadRune()
			if err != nil {
				panic(err)
			}
			_, err = conn.Write([]byte(string(input)))

			if err != nil {
				panic(err)
			}
		}
	}()
	for {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			panic(err)
		}
	}
}
