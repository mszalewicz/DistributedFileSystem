package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	tcpTransport := NewTCPTransport(":42069")
	assert.Equal(t, tcpTransport.listenAddress, ":42069")
}
