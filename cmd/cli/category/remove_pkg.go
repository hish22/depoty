package category

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func RemovePkgInCateCommand() {
	var RemovePkgInCateCommand = &cobra.Command{
		Use:   "remove",
		Short: "Remove a specific package from a catagory.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty remove <package> <category>")
		},
	}
	core.RootCommand.AddCommand(RemovePkgInCateCommand)
}
