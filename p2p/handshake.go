package p2p

import "errors"

// ErrorInvalidHandshake is returned if handshake b/w
// local and remote isn't established.
var ErrorInvalidHandshake = errors.New("invalid handshake")

// handshake func is
type HandShakeFunc func(Peer) error

func HandshakeFn(Peer) error {
	return nil
}
