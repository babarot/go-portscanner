package main

import (
	"flag"
	"os"
	"time"

	portscanner "github.com/b4b4r07/go-portscaner"
)

var (
	host = flag.String("h", "localhost", "host name")
	port = flag.Int("p", 80, "port")
)

func main() {
	flag.Parse()
	ps := portscanner.NewPortScanner(*host, 150*time.Millisecond)
	open := ps.Scan(*port)
	if !open {
		os.Exit(1)
	}
}
