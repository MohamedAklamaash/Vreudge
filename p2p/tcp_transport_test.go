package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenaddr := ":4000"
	trops := TCPTransportOps{ListenAddress: listenaddr, HandshakeFn: HandshakeFn}
	tr := NewTCPTransport(trops)
	assert.Equal(t, tr.ListenAddress, listenaddr) // expecting for equal vals b/w listner addr sent and var assigned

	assert.Nil(t, tr.ListenAndAccept()) // expecting a nil from listenandaccept function
	select {}
}
