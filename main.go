package main

import (
	"github/Tomas/foreverstore/p2p"
	"log"
)

func main()  {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select{}
}

//44:43