package listener

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	// Create a listener on random port
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}
