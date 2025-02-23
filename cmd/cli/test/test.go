package test

import (
	"depoty/cmd/cli/core"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "testing",
		Run: func(cmd *cobra.Command, args []string) {
			// db := badgers.MainDb("/tmp/badger")
			// db.DropAll()
			// db.Close()
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
