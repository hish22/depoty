package category

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCateCommand() {
	var CreateCateCommand = &cobra.Command{
		Use:   "create",
		Short: "create a catagory.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty create <catagory name>")
		},
	}
	core.RootCommand.AddCommand(CreateCateCommand)
}
