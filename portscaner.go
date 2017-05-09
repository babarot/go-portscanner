package portscanner

import (
	"fmt"
	"net"
	"time"
)

type Port int

var (
	DefaultHostname string = "localhost"
)

type Server struct {
	Hostname string
	Timeout  time.Duration
}

func NewServer(hostname string) *Server {
	return &Server{
		Hostname: hostname,
		Timeout:  3 * time.Second,
	}
}

func (s *Server) Scan(port int) bool {
	server := fmt.Sprintf("%s:%d", s.Hostname, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		return false
	}
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), s.Timeout)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func Available(port int) bool {
	return !NewServer(DefaultHostname).Scan(port)
}

func Get() Port {
	addr, err := net.ResolveTCPAddr("tcp", DefaultHostname+":0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	port := l.Addr().(*net.TCPAddr).Port
	return Port(port)
}

func (p Port) Addr() string {
	return fmt.Sprintf(":%d", p)
}
