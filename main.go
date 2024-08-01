package main

import (
	"github/Tomas/foreverstore/p2p"
	"log"
)

func main()  {
	tcpOpts := p2p.TCPTransportOps {
		ListenAddr: 	":3000",
		HandshakeFunc: 	p2p.NOPHandshakeFunc,
		Decoder: 		p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select{}
}

//44:43