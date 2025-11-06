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

func TestSendMagicPacket_WithCustomAddr(t *testing.T) {
	const testMAC = "AA:BB:CC:DD:EE:FF"

	// 1. Set up a listener on a random port.
	listener, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to set up UDP listener: %v", err)
	}
	defer listener.Close()
	listenerAddr := listener.LocalAddr().String()
	t.Logf("Test listener for custom address running on %s", listenerAddr)

	// 2. Use the convenience function with the custom address.
	// This directly tests the logic used by the --via flag.
	sendErrChan := make(chan error, 1)
	go func() {
		sendErrChan <- SendMagicPacket(testMAC, listenerAddr)
	}()

	// 3. Wait for the packet.
	if err := listener.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatalf("Failed to set read deadline: %v", err)
	}
	receivedBuf := make([]byte, 1024)
	n, _, err := listener.ReadFrom(receivedBuf)
	if err != nil {
		t.Fatalf("Failed to read from UDP listener: %v", err)
	}

	// 4. Verify the results.
	if sendErr := <-sendErrChan; sendErr != nil {
		t.Fatalf("SendMagicPacket failed: %v", sendErr)
	}
	if n != 102 {
		t.Errorf("expected to receive 102 bytes, but got %d", n)
	}
}
