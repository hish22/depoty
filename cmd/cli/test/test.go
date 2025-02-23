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
			// listing.OutdatedList()
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
