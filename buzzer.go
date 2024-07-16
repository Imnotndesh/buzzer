package main

import (
	"buzzer/DB_Worker"
	"buzzer/WoL_Worker"
	"fmt"
	"log"
	"os"
)

func main() {
	switch os.Args[1] {
	case "-s":
		var MAC = os.Args[2]
		var IP = os.Args[3]
		var alias = os.Args[4]
		err := DB_Worker.StoreMachine(alias, MAC, IP)
		if err != nil {
			log.Println("Error storing machine: ", err)
		}
	case "-b":
		var MAC = os.Args[2]
		packet, err := WoL_Worker.CreateMagicPacket(MAC)
		if err != nil {
			log.Println("Error creating magic packet: ", err)
		}
		err = packet.Send()
		if err != nil {
			log.Println("Error sending magic packet: ", err)
		}
		fmt.Println("Packet sent successfully to machine: ", MAC)
	}

}
