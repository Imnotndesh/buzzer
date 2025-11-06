package DB_Worker

import (
	"buzzer/WoL_Worker"
	"fmt"
	"log"
	"strconv"

	"github.com/tidwall/buntdb"
)

//type machine struct {
//	Alias string
//	MAC   string
//}

// DB manages interactions with the buntdb database.
type DB struct {
	conn *buntdb.DB
}

// New opens a database at the given path and returns a new DB instance.
func New(path string) (*DB, error) {
	conn, err := buntdb.Open(path)
	if err != nil {
		return nil, err
	}
	return &DB{conn: conn}, nil
}

// Close closes the database connection.
func (d *DB) Close() error {
	if d.conn != nil {
		return d.conn.Close()
	}
	return nil
}

// StoreMachine saves a new alias and MAC address to the database.
func (d *DB) StoreMachine(alias string, macStr string) error {
	return d.conn.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(alias, macStr, nil)
		return err
	})
}

// GetStoredMac retrieves a MAC address for a given alias.
func (d *DB) GetStoredMac(alias string) (MAC string, err error) {
	var macStr string
	err = d.conn.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(alias)
		if err != nil {
			return err
		}
		macStr = val
		return nil
	})
	if err != nil {
		return "", err
	}
	return macStr, err
}

// WakeWithAlias retrieves a MAC from the DB and sends a WoL packet.
func (d *DB) WakeWithAlias(alias string, customAddr ...string) error {
	mac, err := d.GetStoredMac(alias)
	if err != nil {
		return fmt.Errorf("unable to get stored mac for alias %q: %w", alias, err)
	}
	return WoL_Worker.SendMagicPacket(mac, customAddr...)
}

// EditMachineDetails updates the MAC address for an existing alias.
func (d *DB) EditMachineDetails(alias string, macStr string) error {
	return d.StoreMachine(alias, macStr)
}

// ListAllMachines iterates through all entries and prints them.
func (d *DB) ListAllMachines() error {
	return d.conn.View(func(tx *buntdb.Tx) error {
		index := 0
		err := tx.Ascend("", func(key, value string) bool {
			fmt.Println(strconv.Itoa(index) + "----------\n" + "Alias : " + key + "\n" + "MAC_ADDRESS: " + value)
			index++
			return true
		})
		return err
	})
}

// DeleteEntry removes an entry from the database by its alias.
func (d *DB) DeleteEntry(alias string) error {
	err := d.conn.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(alias)
		return err
	})
	if err != nil {
		return fmt.Errorf("cannot delete pair for alias %q: %w", alias, err)
	}
	log.Println("Entry removed successfully")
	return nil
}

// ListAllMachineAliases iterates and prints only the alias keys, one per line.
func (d *DB) ListAllMachineAliases() error {
	return d.conn.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(key, value string) bool {
			fmt.Println(key)
			return true
		})
	})
}
