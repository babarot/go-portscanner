package portscanner

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type Port int

var (
	DefaultHostname string = "localhost"
)

type Server struct {
	Hostname string
	Timeout  time.Duration
	Port     Port
}

func NewServer(hostname string) *Server {
	return &Server{
		Hostname: hostname,
		Timeout:  3 * time.Second,
		Port:     Get(),
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

func getPort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", DefaultHostname+":0")
	if err != nil {
		return
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

func Get() Port {
	p, _ := getPort()
	return Port(p)
}

func GetWith(port int) Port {
	if Available(port) {
		return Port(port)
	}
	return Get()
}

func (p Port) Addr() string {
	return fmt.Sprintf(":%d", p)
}

func (p Port) Listen() error {
	return http.ListenAndServe(p.Addr(), http.DefaultServeMux)
}
