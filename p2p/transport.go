package p2p

// Peer is an interface that represent a remote node
type Peer interface {
	Close() error
}

// Transport handles communication between nodes.
// Can be of the form TCP, UDP, WebSockets, etc.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
