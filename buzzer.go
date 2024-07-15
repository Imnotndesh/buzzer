package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var MAC string
	var packet []byte
	MAC = os.Args[1]
	IP := net.ParseIP(os.Args[2])
	mac, err := net.ParseMAC(MAC)
	if err != nil {
		log.Println("Error parsing MAC")
		os.Exit(1)
	}
	packet = make([]byte, 102)
	for i := 0; i < 6; i++ {
		packet[i] = 0xff
	}
	copy(packet[6:], append(mac, mac...))
	conn, err := net.ListenPacket("udp", "")
	if err != nil {
		log.Println("Error creating socket", err)
	}
	defer conn.Close()
	_, err = conn.WriteTo(packet, &net.UDPAddr{IP: IP, Port: 9})
	if err != nil {
		log.Println("Error sending packet!", err)
	}
	fmt.Println("Started machine!")
}
