package main

import (
	"fmt"
	"log"

	"github.com/kaiqiz/distributedfilesystem/p2p"
)

func main() {
	// TCP server
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.GOBDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccpet(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("We Gucci!")

	select {}
}
