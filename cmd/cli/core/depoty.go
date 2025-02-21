package core

import (
	"depoty/cmd/tui"
	"depoty/internal/badger"
	"log"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "depoty",
	Short: "Root command for managing Depoty.",
	Long:  "Depoty is an advanced extension of package managers, offering enhanced local management features for package handling and control.",
	Run: func(cmd *cobra.Command, args []string) {

		db := badger.MainDb()

		defer db.Close()

		if _, err := badger.Read(db, []byte("initDone")); err != nil {
			log.Fatal("Please Start the Configuration process before using the TUI, type > Depoty init")
		}

		tui.TuiStart()

	},
}
