package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer reprs a remote node over a tcp established connection
type TCPPeer struct {
	conn net.Conn // underlying conn of the peer

	// if we dial and retrive a conn => outbound
	// if we accept and retrive a conn => inbound i.e outbound == false
	outbound bool // if we send req to the peer to accept files from us,it's outbound
}

type TCPTransportOps struct {
	ListenAddress string
	HandshakeFn   HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOps
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	// check if we are able to make a network connection to that machine
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop() // create a go routine cuz for handling independent processes in our software
	return nil
}

func (t *TCPTransport) startAcceptLoop() error {
	for {
		conn, err := t.listener.Accept() // it returns conn string or error occured
		if err != nil {
			fmt.Println("error in accepting the connection loop:", err)
			return err
		}
		fmt.Printf("new incoming connection:%+v\n", conn)
		go t.handleConnection(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true) // each new connection is a p2p conn
	if err := t.HandshakeFn(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error:%s\n", err)
		return
	}
	msg := &Message{}
	// buffer := make([]byte, 2000) // creates an empty buffer so that msg can be saved there
	for {
		// n, err := conn.Read(buffer) // reading from the buffer
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp decode error:%s\n", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("message:%+v\n", msg)
	}
}
