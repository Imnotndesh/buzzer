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
		var alias = os.Args[2]
		var MAC = os.Args[3]
		err := DB_Worker.StoreMachine(alias, MAC)
		if err != nil {
			log.Println("Error storing machine: ", err)
		}
		fmt.Println("Machine stored successfully")
	case "-w":
		var alias = os.Args[2]
		err := DB_Worker.WakeWithAlias(alias)
		if err != nil {
			log.Println("Error waking "+alias+": ", err)
		}
		fmt.Println("Waking " + alias + " ...")
	case "-g":
		var alias = os.Args[2]
		Mac, err := DB_Worker.GetStoredMac(alias)
		if err != nil {
			log.Println("Error getting stored mac: ", err)
		}
		fmt.Println(alias + " is tied to: " + Mac)
		// TODO: ADD OPTION TO SAVED MAC ADDRESS BY PASSING A MAC ADDRESS
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
		fmt.Println("Waking" + MAC + " ...")
	}

}
