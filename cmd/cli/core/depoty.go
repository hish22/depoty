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
	Short: "Root command for managing Depoty.",
	Long:  "Depoty is an advanced extension of package managers, offering enhanced local management features for package handling and control.",
	Run: func(cmd *cobra.Command, args []string) {
		// Open badger
		dbConfig := badgers.MainDb("/tmp/badger/config")

		fmt.Println("Checking Initialization Values..")
		// Check if init process is done
		if _, err := badgers.Read(dbConfig, []byte("initDone")); err != nil {
			log.Fatal("Please Start the Configuration process before using the TUI, type > Depoty init")
		}

		dbConfig.Close()

		fmt.Println("Fetching Outdated Packages..")

		// Check Outdated Packages
		listing.OutdatedList()

		fmt.Println("Fetching Packages list...")

		// Start the TUI app
		tui.TuiStart()

	},
}
