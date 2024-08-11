package p2p

import "net"

// p2p msg holds any data to and from each transport
// b/w 2 nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
