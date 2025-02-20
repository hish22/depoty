package list

import (
	"depoty/cmd/cli/core"
	"depoty/internal/badger"
	"fmt"

	"github.com/spf13/cobra"
)

func ListCommand() {
	var ListCommand = &cobra.Command{
		Use:   "list",
		Short: "Displays all monitored packages.",
		Run: func(cmd *cobra.Command, args []string) {
			db := badger.MainDb()
			fmt.Println(string(badger.Read(db, []byte("name"))))
			defer db.Close()
		},
	}
	core.RootCommand.AddCommand(ListCommand)
}
