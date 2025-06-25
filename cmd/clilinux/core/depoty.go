//go:build linux

package core

import (
	"depoty/cmd/tui"
	"depoty/internal/badgers"
	"depoty/internal/listing"
	"depoty/internal/util/common"
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
		dbConfig := badgers.MainDb("/system/apt/config")

		fmt.Println("Checking Initialization Values..")
		// Check if init process is done
		if _, err := badgers.Read(dbConfig, []byte("initDone")); err != nil {
			log.Fatal("Please Start the Configuration process before using the TUI, type > depoty init")
		}

		dbConfig.Close()

		fmt.Println("Fetching Outdated Packages..")

		// Check Outdated Packages
		listing.OutdatedList()

		// Update apt packages
		fmt.Println("Update apt package list => apt update")
		common.ExecutePrevScript("apt update", "")

		fmt.Println("Fetching Packages list...")

		// Start the TUI app
		fmt.Println("Starting Depoty..")

		tui.TuiStart()

	},
}
