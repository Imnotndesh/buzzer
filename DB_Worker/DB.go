package DB_Worker

import (
	"errors"
	"fmt"
	"net"
)

type machine struct {
	Alias string
	IP    net.IP
	MAC   net.HardwareAddr
}

func StoreMachine(alias string, MACstr string, IP_Address string) error {
	MAC, err := net.ParseMAC(MACstr)
	if err != nil {
		return errors.New("Failed to parse MAC")
	}
	IP := net.ParseIP(IP_Address)
	newMachine := machine{alias, IP, MAC}
	fmt.Println("Machine " + newMachine.Alias + " Stored with details { " + newMachine.MAC.String() + newMachine.IP.String() + " }")
	return nil
}
