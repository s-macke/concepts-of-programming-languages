package main

import (
	"bufio"
	"fmt"
	"io"
	_ "io"
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

type Session struct {
	stdin  io.WriteCloser
	stdout io.ReadCloser
	cmd    *exec.Cmd
	reader *bufio.Reader
}

func (s *Session) Read(_ *struct{}, reply *rune) error {
	r, _, err := s.reader.ReadRune()
	*reply = r
	return err
}

func (s *Session) Write(input *rune, _ *struct{}) error {
	_, err := s.stdin.Write([]byte(string(*input)))
	return err
}

func NewSession() *Session {
	var s = new(Session)
	var err error
	s.cmd = exec.Command("./frotz/dfrotz", "-p", "-m", "frotz/905.z5")
	s.stdin, err = s.cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	s.stdout, err = s.cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}

	s.reader = bufio.NewReader(s.stdout)

	if err := s.cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", s.cmd.Path, err.Error())
	}
	return s
}

func RunRPCSession(c net.Conn) {
	server := rpc.NewServer()
	var s *Session
	s = NewSession()
	defer func() {
		s.cmd.Process.Kill()
		s.cmd.Wait()
	}()
	server.Register(s)
	server.ServeConn(c)
	fmt.Println("Closing connection from ", c.RemoteAddr())
}

func Listen() {
	l, e := net.Listen("tcp", ":1234")
	defer l.Close()
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		c, err := l.Accept()
		log.Printf("request from %v\n", c.RemoteAddr())
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go RunRPCSession(c)
	}
}

func main() {
	Listen()
}
