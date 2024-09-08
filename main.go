package main

import (
	"bytes"
	"fmt"
	//"fmt"
	"github/Tomas/foreverstore/p2p"
	//"io/ioutil"
	"log"
	"time"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	
	tcptransportOpts := p2p.TCPTransportOpts {
		ListenAddr: 	listenAddr,
		HandshakeFunc: 	p2p.NOPHandshakeFunc,
		Decoder: 		p2p.DefaultDecoder{},
	}
	
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)
	
	fileServerOpts := FileServerOpts {
		StorageRoot: 		listenAddr + "_network",
		PathTransformFunc: 	CASPathTransformFunc,
		Transport: 			tcpTransport,
		BootstrapNodes:		nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	
	go func() {
		log.Fatal(s1.Start())
	}()

	time.Sleep(2 * time.Second)

	go s2.Start()
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		data := bytes.NewReader([]byte("my big data file here!"))
		s2.StoreData(fmt.Sprintf("myprivatedata_%d", i) , data)
		time.Sleep(1 * time.Millisecond)
	}



	/* r, err := s1.Get("myprivatedata")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b)) */
 
	select{}
}

//6:38:54
