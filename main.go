package main

import (
	"fmt"
	"log"

	"github.com/kaiqiz/distributedfilesystem/p2p"
)

func main() {
	// TCP server
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccpet(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("We Gucci!")
	select {}
}
