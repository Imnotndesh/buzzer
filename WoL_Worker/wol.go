package WoL_Worker

import (
	"fmt"
	"log"
	"net"
)

// Packet represents a standard Wake-on-LAN magic packet.
// It consists of 6 bytes of 0xFF followed by 16 repetitions of the target MAC address.
type Packet [102]byte

// Client is a Wake-on-LAN client that can send magic packets.
type Client struct {
	addr string
}

// NewClient creates a new Wake-on-LAN client.
// By default, it's configured to send to the global broadcast address "255.255.255.255" on port "9".
func NewClient() *Client {
	return &Client{
		addr: "255.255.255.255:9",
	}
}

// WithBroadcastAddr allows setting a custom broadcast address and port (e.g., "192.168.1.255:7").
func (c *Client) WithBroadcastAddr(addr string) *Client {
	c.addr = addr
	return c
}

// CreateMagicPacket generates a magic packet for the given MAC address string.
func CreateMagicPacket(macStr string) (Packet, error) {
	var magicPacket Packet

	MAC, err := net.ParseMAC(macStr)
	if err != nil {
		return magicPacket, fmt.Errorf("error parsing MAC address %q: %w", macStr, err)
	}

	// The magic packet starts with 6 bytes of 0xFF.
	for i := 0; i < 6; i++ {
		magicPacket[i] = 0xFF
	}

	// It is followed by 16 repetitions of the MAC address.
	offset := 6
	for i := 0; i < 16; i++ {
		copy(magicPacket[offset:], MAC)
		offset += len(MAC)
	}

	return magicPacket, nil
}

// Send transmits the magic packet.
func (c *Client) Send(p Packet) error {
	conn, err := net.Dial("udp", c.addr)
	if err != nil {
		return fmt.Errorf("failed to dial UDP for WoL packet: %w", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing WoL connection: %v\n", err)
		}
	}(conn)

	_, err = conn.Write(p[:])
	if err != nil {
		return fmt.Errorf("failed to write WoL packet: %w", err)
	}
	return nil
}

// SendMagicPacket is a convenience function that creates and sends a magic packet
// to the default broadcast address.
func SendMagicPacket(macStr string, customAddr ...string) error {
	packet, err := CreateMagicPacket(macStr)
	if err != nil {
		return err
	}
	client := NewClient()
	if len(customAddr) > 0 && customAddr[0] != "" {
		client.WithBroadcastAddr(customAddr[0])
	}
	return client.Send(packet)
}
