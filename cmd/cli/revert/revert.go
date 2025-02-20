package revert

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func RevertCommand() {
	var RevertCommand = &cobra.Command{
		Use:   "revert",
		Short: "Rolls back a package to a specified version.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty revert <Package> <version>")
		},
	}
	core.RootCommand.AddCommand(RevertCommand)
}
