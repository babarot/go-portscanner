package main

import (
	"fmt"
	"log"

	"github.com/b4b4r07/go-portscanner"
)

func main() {
	fmt.Printf("%#v\n", portscanner.Get())
	fmt.Printf("%#v\n", portscanner.Get().Addr())

	server := portscanner.NewServer("localhost")
	fmt.Printf("%#v\n", server.Scan(8000))

	// Alias for portscaner.Server.Scan()
	fmt.Printf("%#v\n", portscanner.Available(8000))

	log.Println(portscanner.Get().Listen())
}
