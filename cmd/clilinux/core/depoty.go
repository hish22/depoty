//go:build linux

package core

import (
	"depoty/cmd/tui"
	"depoty/internal/badgers"
	"depoty/internal/listing"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "depoty",
	Short: "Root command to start the TUI app.",
	Long:  "Depoty is an extension to Chocolatey package manager, offering enhanced local management features for package handling and control.",
	Run: func(cmd *cobra.Command, args []string) {

		// Open badger
		dbConfig := badgers.MainDb("/system/choco/config")

		fmt.Println("Checking Initialization Values..")
		// Check if init process is done
		if _, err := badgers.Read(dbConfig, []byte("initDone")); err != nil {
			log.Fatal("Please Start the Configuration process before using the TUI, type > depoty init")
		}

		dbConfig.Close()

		fmt.Println("Fetching Outdated Packages..")

		// Check Outdated Packages
		listing.OutdatedList()

		fmt.Println("Fetching Packages list...")

		// Start the TUI app
		fmt.Println("Starting Depoty..")

		tui.TuiStart()
		fmt.Println("Soon with apt package manager")

	},
}
