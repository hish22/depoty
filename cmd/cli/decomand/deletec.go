package decomand

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func DeleteCommand() {
	var DeleteCommand = &cobra.Command{
		Use:   "delete",
		Short: "Removes specified packages.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty delete <Packages>")
		},
	}
	core.RootCommand.AddCommand(DeleteCommand)
}
