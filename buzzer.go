package main

import (
	"buzzer/DB_Worker"
	"buzzer/WoL_Worker"
	"fmt"
	"log"
	"os"
)

const (
	version = "0.2.2"
)

func helpMessage() string {
	helpMessage := `
	Usage: 
		buzzer [options]
	
	Options:
		-b [MAC_ADDRESS]						Wakes machine using the passed MAC ADDRESS
		-e [ALIAS] [MAC_ADDRESS]					Changes MAC_ADDRESS value of passed ALIAS to passed MAC_ADDRESS 
		-g [STORED_ALIAS]							Fetches MAC ADDRESS bound to the passed ALIAS
		-h								Help text
		-l								Prints out all stored aliases along with their MAC_ADDRESSES
		-s [ALIAS] [MAC_ADDRESS]					Binds the passed alias and saves it
		-V								Prints version of the program
		-w [ALIAS]							Wakes machine using the passed ALIAS
`
	return helpMessage
}
func main() {
	if len(os.Args)-1 == 0 {
		fmt.Println(helpMessage())
	} else {
		switch os.Args[1] {
		case "-s":
			var alias = os.Args[2]
			var MAC = os.Args[3]
			err := DB_Worker.StoreMachine(alias, MAC)
			if err != nil {
				log.Println("Error storing machine: ", err)
				os.Exit(1)
			}
			fmt.Println("Machine stored successfully")
		case "-e":
			var alias = os.Args[2]
			var newMAC = os.Args[3]
			err := DB_Worker.EditMachineDetails(alias, newMAC)
			if err != nil {
				log.Println("Error editing machine: ", err)
				os.Exit(1)
			}
			fmt.Println("Machine edited successfully")

		case "-w":
			var alias = os.Args[2]
			err := DB_Worker.WakeWithAlias(alias)
			if err != nil {
				log.Println("Error waking "+alias+": ", err)
				os.Exit(1)
			}
			fmt.Println("Waking " + alias + " ...")
		case "-g":
			var alias = os.Args[2]
			Mac, err := DB_Worker.GetStoredMac(alias)
			if err != nil {
				log.Println("Error getting stored mac: ", err)
				os.Exit(1)
			}
			fmt.Println(alias + " is tied to: " + Mac)
		case "-b":
			var MAC = os.Args[2]
			packet, err := WoL_Worker.CreateMagicPacket(MAC)
			if err != nil {
				log.Println("Error creating magic packet: ", err)
				os.Exit(1)
			}
			err = packet.Send()
			if err != nil {
				log.Println("Error sending magic packet: ", err)
				os.Exit(1)
			}
			fmt.Println("Waking" + MAC + " ...")
		case "-h":
			fmt.Println(helpMessage())
		case "-l":
			err := DB_Worker.ListAllMachies()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		case "-v":
			fmt.Println(version)
		}
	}
}
