package p2p

import (
	"log/slog"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	peersMutex sync.RWMutex
	peers      map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() (err error) {
	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		connection, err := t.listener.Accept()
		if err != nil {
			slog.Error("Error during accepting TCP connetion", "err", err)
		}

	}
}
