package main

import (
	"log"

	"github.com/MohamedAklamaash/Vreudge/p2p"
)

func main() {
	tcpOps := p2p.TCPTransportOps{
		ListenAddress: ":3000",
		HandshakeFn:   p2p.HandshakeFn,
		Decoder:       &p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOps)
	// listen and accept listens for some conn in the
	// specificed listen addr
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
