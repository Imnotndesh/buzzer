package main

import (
	"buzzer/DB_Worker"
	"buzzer/WoL_Worker"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const (
	version = "v1.2.0"
)

func getDBPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("could not get user config directory: %w", err)
	}
	buzzerDir := filepath.Join(configDir, "buzzer")
	if err := os.MkdirAll(buzzerDir, 0755); err != nil {
		return "", fmt.Errorf("could not create config directory %q: %w", buzzerDir, err)
	}
	return filepath.Join(buzzerDir, ".machines"), nil
}

func main() {
	dbpath, err := getDBPath()
	if err != nil {
		log.Fatalf("Could not get current data path")
	}
	db, err := DB_Worker.New(dbpath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func(db *DB_Worker.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	app := &cli.App{
		Name:    "buzzer",
		Usage:   "A simple Wake-on-LAN (WoL) command-line tool",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:      "store",
				Aliases:   []string{"s"},
				Usage:     "Saves or stores a MAC address with a memorable alias",
				ArgsUsage: "[ALIAS] [MAC_ADDRESS]",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return cli.ShowSubcommandHelp(c)
					}
					alias := c.Args().Get(0)
					mac := c.Args().Get(1)
					if err := db.StoreMachine(alias, mac); err != nil {
						return cli.Exit(fmt.Sprintf("Error storing machine: %v", err), 1)
					}
					fmt.Println("Machine stored successfully")
					return nil
				},
			},
			{
				Name:      "edit",
				Aliases:   []string{"e"},
				Usage:     "Edits an existing entry to assign a new MAC address to an alias",
				ArgsUsage: "[ALIAS] [NEW_MAC_ADDRESS]",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return cli.ShowSubcommandHelp(c)
					}
					alias := c.Args().Get(0)
					newMAC := c.Args().Get(1)
					if err := db.EditMachineDetails(alias, newMAC); err != nil {
						return cli.Exit(fmt.Sprintf("Error editing machine: %v", err), 1)
					}
					fmt.Println("Machine edited successfully")
					return nil
				},
			},
			{
				Name:      "wake",
				Aliases:   []string{"w"},
				Usage:     "Wakes a machine using its stored alias",
				ArgsUsage: "[ALIAS]",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.ShowSubcommandHelp(c)
					}
					alias := c.Args().First()
					if err := db.WakeWithAlias(alias); err != nil {
						return cli.Exit(fmt.Sprintf("Error waking %q: %v", alias, err), 1)
					}
					fmt.Printf("Waking %s ...\n", alias)
					return nil
				},
			},
			{
				Name:      "get",
				Aliases:   []string{"g"},
				Usage:     "Gets and displays the MAC address associated with a stored alias",
				ArgsUsage: "[ALIAS]",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.ShowSubcommandHelp(c)
					}
					alias := c.Args().First()
					mac, err := db.GetStoredMac(alias)
					if err != nil {
						return cli.Exit(fmt.Sprintf("Error getting stored mac: %v", err), 1)
					}
					fmt.Printf("%s is tied to: %s\n", alias, mac)
					return nil
				},
			},
			{
				Name:      "broadcast",
				Aliases:   []string{"b"},
				Usage:     "Wakes a machine directly via its MAC address (Broadcast)",
				ArgsUsage: "[MAC_ADDRESS]",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.ShowSubcommandHelp(c)
					}
					mac := c.Args().First()
					if err := WoL_Worker.SendMagicPacket(mac); err != nil {
						return cli.Exit(fmt.Sprintf("Error sending magic packet: %v", err), 1)
					}
					fmt.Printf("Waking %s ...\n", mac)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Lists all stored aliases and their corresponding MAC addresses",
				Action: func(c *cli.Context) error {
					return db.ListAllMachines()
				},
			},
			{
				Name:      "remove",
				Aliases:   []string{"r"},
				Usage:     "Removes an alias and its MAC address from the database",
				ArgsUsage: "[ALIAS]",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.ShowSubcommandHelp(c)
					}
					alias := c.Args().First()
					return db.DeleteEntry(alias)
				},
			},
			{
				Name:   "list-raw",
				Hidden: true,
				Action: func(c *cli.Context) error {
					return db.ListAllMachineAliases()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
