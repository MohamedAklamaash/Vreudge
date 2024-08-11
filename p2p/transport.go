package p2p

// peer reprs the remote node
type Peer interface {
}

// transport is anything that can handle the communication b/w nodes in the network
// it can handle tcp,udp,socket,...
type Transport interface {
	ListenAndAccept() error
}
