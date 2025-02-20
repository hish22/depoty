package update

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func UpdateCommand() {
	var UpdateCommand = &cobra.Command{
		Use:   "update",
		Short: "Updates specified packages to the latest version.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty update <Packages>")
		},
	}
	core.RootCommand.AddCommand(UpdateCommand)
}
