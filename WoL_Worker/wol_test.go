package WoL_Worker

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestSendMagicPacket_Integration(t *testing.T) {
	const testMAC = "0A:0B:0C:0D:0E:0F"

	listener, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to set up UDP listener: %v", err)
	}
	defer listener.Close()
	listenerAddr := listener.LocalAddr().String()
	t.Logf("Test listener running on %s", listenerAddr)
	expectedPacket, err := CreateMagicPacket(testMAC)
	if err != nil {
		t.Fatalf("Failed to create expected magic packet: %v", err)
	}
	client := NewClient().WithBroadcastAddr(listenerAddr)
	sendErrChan := make(chan error, 1)
	go func() {
		t.Log("Sending magic packet...")
		sendErrChan <- client.Send(expectedPacket)
	}()
	if err := listener.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatalf("Failed to set read deadline on listener: %v", err)
	}

	receivedBuf := make([]byte, 1024)
	n, addr, err := listener.ReadFrom(receivedBuf)
	if err != nil {
		t.Fatalf("Failed to read from UDP listener: %v", err)
	}

	t.Logf("Received %d bytes from %s", n, addr)

	if sendErr := <-sendErrChan; sendErr != nil {
		t.Fatalf("client.Send failed: %v", sendErr)
	}
	if n != len(expectedPacket) {
		t.Errorf("expected to receive %d bytes, but got %d", len(expectedPacket), n)
	}
	receivedPacket := receivedBuf[:n]
	if !bytes.Equal(expectedPacket[:], receivedPacket) {
		t.Errorf("received packet does not match expected magic packet")
		t.Logf("Expected: %x", expectedPacket)
		t.Logf("Received: %x", receivedPacket)
	}
}
