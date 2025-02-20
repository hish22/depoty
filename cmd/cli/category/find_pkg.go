package category

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func FingPkgInCateCommand() {
	var FingPkgCommand = &cobra.Command{
		Use:   "find",
		Short: "Check a package in specific catagory.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty drop <catagory>")
		},
	}
	core.RootCommand.AddCommand(FingPkgCommand)
}
