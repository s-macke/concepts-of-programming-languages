package main

import (
	_ "bufio"
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
}

func (s *Session) Read(_ *struct{}, reply *[]byte) error {
	p := make([]byte, 1)
	_, err := s.stdout.Read(p)
	*reply = p
	return err
}

func (s *Session) Write(input *[]byte, _ *struct{}) error {
	_, err := s.stdin.Write(*input)
	return err
}

func NewSession() *Session {
	var s = new(Session)
	var err error
	s.cmd = exec.Command("./frotz/dfrotz", "-p", "-m", "905.z5")
	s.stdin, err = s.cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	s.stdout, err = s.cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	if err := s.cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", s.cmd.Path, err.Error())
	}
	return s
}

func RunRPCSession(c net.Conn) {
	server := rpc.NewServer()
	var s *Session
	s = NewSession()
	defer s.cmd.Process.Kill()
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
		fmt.Printf("request from %v\n", c.RemoteAddr())
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
