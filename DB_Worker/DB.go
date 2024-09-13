package DB_Worker

import (
	"buzzer/WoL_Worker"
	"errors"
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"strconv"
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
	dbName string = ".machines"
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
func EditMachineDetails(alias string, MACstr string) error {
	newMachine := machine{alias, MACstr}
	db, err := initializeDB(dbName)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(newMachine.Alias)
		if err != nil {
			return err
		}
		return nil
	})
	err = StoreMachine(newMachine.Alias, newMachine.MAC)
	if err != nil {
		return err
	}
	return nil
}
func ListAllMachines() error {
	db, err := initializeDB(dbName)
	if err != nil {
		return err
	}
	err = db.View(func(tx *buntdb.Tx) error {
		var index int = 0
		err = tx.Ascend("", func(key, value string) bool {
			fmt.Println(strconv.Itoa(index) + "----------\n" + "Alias : " + key + "\n" + "MAC_ADDRESS: " + value)
			index += 1
			return true
		})
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
func DeleteEntry(alias string) error {
	db, err := initializeDB(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *buntdb.Tx) error {
		_, err = tx.Delete(alias)
		return err
	})
	if err != nil {
		log.Fatal("Cannot delete pair: ", err)
	}
	log.Println("Entry removed successfully")
	return nil
}
