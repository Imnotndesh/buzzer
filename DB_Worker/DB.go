package DB_Worker

import (
	"buzzer/WoL_Worker"
	"errors"
	"fmt"
	"github.com/tidwall/buntdb"
)

type machine struct {
	Alias string
	MAC   string
}

func initializeDB(DBname string) (db *buntdb.DB, err error) {
	db, err = buntdb.Open(DBname)
	if err != nil {
		return nil, err
	}
	return
}

const (
	dbName string = "./DB_Worker/.machines"
)

func StoreMachine(alias string, MACstr string) error {
	newMachine := machine{alias, MACstr}
	fmt.Println("Machine " + newMachine.Alias + " Stored with details { " + newMachine.MAC + " }")
	db, err := initializeDB(dbName)
	defer func(db *buntdb.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	if err != nil {
		return errors.New("failed to initialize DB")
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(newMachine.Alias, newMachine.MAC, nil)
		return err
	})
	return nil
}
func GetStoredMac(alias string) (MAC string, err error) {
	var MACstr string
	db, err := initializeDB(dbName)
	if err != nil {
		return "", err
	}
	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(alias)
		if err != nil {
			return err
		}
		MACstr = val
		return nil
	})
	if err != nil {
		return "", err
	}
	return MACstr, err
}
func WakeWithAlias(alias string) error {
	Mac, err := GetStoredMac(alias)
	if err != nil {
		return errors.New("unable to get stored mac")
	}
	packet, err := WoL_Worker.CreateMagicPacket(Mac)
	if err != nil {
		return errors.New("unable to create packet")
	}
	err = packet.Send()
	if err != nil {
		return errors.New("unable to send packet")
	}
	return nil
}
