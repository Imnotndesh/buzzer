package main

import (
	"buzzer/DB_Worker"
	"buzzer/WoL_Worker"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	version = "v1.0.0"
)

func helpMessage() string {
	helpMessage := `
	Usage: 
		buzzer [options]
	
	Options:
		-B [MAC_ADDRESS]						Wakes machine using the passed MAC ADDRESS
		-E [ALIAS] [MAC_ADDRESS]					Changes MAC_ADDRESS value of passed ALIAS to passed MAC_ADDRESS 
		-G [STORED_ALIAS]							Fetches MAC ADDRESS bound to the passed ALIAS
		-R [STORED_ALIAS]							Removes the entire entry from database
		-H								Help text
		-L								Prints out all stored aliases along with their MAC_ADDRESSES
		-S [ALIAS] [MAC_ADDRESS]					Binds the passed alias and saves it
		-V								Prints version of the program
		-W [ALIAS]							Wakes machine using the passed ALIAS
`
	return helpMessage
}
func main() {
	if len(os.Args)-1 == 0 {
		fmt.Println(helpMessage())
	} else {
		switch strings.ToUpper(os.Args[1]) {
		case "-S":
			var alias = os.Args[2]
			var MAC = os.Args[3]
			err := DB_Worker.StoreMachine(alias, MAC)
			if err != nil {
				log.Println("Error storing machine: ", err)
				os.Exit(1)
			}
			fmt.Println("Machine stored successfully")
		case "-E":
			var alias = os.Args[2]
			var newMAC = os.Args[3]
			err := DB_Worker.EditMachineDetails(alias, newMAC)
			if err != nil {
				log.Println("Error editing machine: ", err)
				os.Exit(1)
			}
			fmt.Println("Machine edited successfully")

		case "-W":
			var alias = os.Args[2]
			err := DB_Worker.WakeWithAlias(alias)
			if err != nil {
				log.Println("Error waking "+alias+": ", err)
				os.Exit(1)
			}
			fmt.Println("Waking " + alias + " ...")
		case "-G":
			var alias = os.Args[2]
			Mac, err := DB_Worker.GetStoredMac(alias)
			if err != nil {
				log.Println("Error getting stored mac: ", err)
				os.Exit(1)
			}
			fmt.Println(alias + " is tied to: " + Mac)
		case "-B":
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
		default:
			fallthrough
		case "-R":
			alias := os.Args[2]
			err := DB_Worker.DeleteEntry(alias)
			if err != nil {
				fmt.Println("Error deleting entry: ", err)
			}
		case "-H":
			fmt.Println(helpMessage())
		case "-L":
			err := DB_Worker.ListAllMachines()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		case "-V":
			fmt.Println(version)
		}
	}
}
