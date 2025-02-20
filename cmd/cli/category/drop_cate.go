package category

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func DropCateCommand() {
	var DropCateCommand = &cobra.Command{
		Use:   "drop",
		Short: "Deletes the entire catagory.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty drop <catagory>")
		},
	}
	core.RootCommand.AddCommand(DropCateCommand)
}
