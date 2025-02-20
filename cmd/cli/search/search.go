package search

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func SearchCommand() {
	var SearchCommand = &cobra.Command{
		Use:   "search",
		Short: "find a specific package.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty search <package>")
		},
	}
	core.RootCommand.AddCommand(SearchCommand)
}
