package portscanner

import (
	"fmt"
	"net"
	"time"
)

type PortScanner struct {
	host    string
	timeout time.Duration
}

func NewPortScanner(host string, timeout time.Duration) *PortScanner {
	return &PortScanner{
		host,
		timeout,
	}
}

func (h PortScanner) getHostPort(port int) string {
	return fmt.Sprintf("%s:%d", h.host, port)
}

func (h PortScanner) GetOpenedPorts(startPort int, Endport int) []int {
	ret := []int{}
	for port := startPort; port <= Endport; port++ {
		if h.Scan(port) {
			ret = append(ret, port)
		}
	}
	return ret
}

func (h PortScanner) Scan(port int) bool {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", h.getHostPort(port))
	if err != nil {
		return false
	}
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), h.timeout)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
