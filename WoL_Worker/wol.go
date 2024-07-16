package WoL_Worker

import (
	"errors"
	"log"
	"net"
)

type Packet [102]byte

func CreateMagicPacket(MACstr string) (Packet, error) {
	var stream = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	var MagicPacket Packet
	MAC, err := net.ParseMAC(MACstr)
	if err != nil {
		return MagicPacket, errors.New("error while parsing MAC")
	}
	copy(MagicPacket[0:], stream)
	macLen := len(MAC)
	for i := 0; i < 16; i++ {
		copy(MagicPacket[macLen:], MAC)
		macLen += len(MAC)
	}
	return MagicPacket, nil
}
func (p *Packet) Send() error {
	conn, err := net.Dial("udp", "255.255.255.255:7")
	if err != nil {
		return err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection")
		}
	}(conn)
	_, err = conn.Write(p[:])
	return err
}
