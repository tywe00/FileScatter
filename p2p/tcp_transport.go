package p2p

import (
	"fmt"
	"net"
	"sync"
)

//TCPPeer represents the remote node over a TCP established connection.
type TCPPeer struct {

	//conn is the underlying connection of the peer
	conn		net.Conn

	// if we dial and retrieve a conn => outbound == true
	// if we accept and retrieve a conn => outbound == false
	outbound 	bool 

}

type TCPTransportOps struct {
	ListenAddr		string
	HandshakeFunc 	HandshakeFunc
	Decoder 		Decoder
}

type TCPTransport struct {
	TCPTransportOps
	listener 		net.Listener
	shakeHands		HandshakeFunc
	decoder 		Decoder

	mu				sync.RWMutex
	peers 			map[net.Addr]Peer
}

func newTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:		conn, 
		outbound: 	outbound,
	}
}



func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error 

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for{
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConn(conn)
	}

}

type Temp struct {}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := newTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}
	
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		fmt.Printf("message: %+v\n", msg)
	}

}

