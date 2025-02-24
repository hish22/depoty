package test

import (
	"depoty/cmd/cli/core"
	"depoty/internal/deletion"
	"fmt"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "testing",
		Run: func(cmd *cobra.Command, args []string) {
			var pkgs []string = []string{"php", "notion"}
			deletion.DropAllPkgs(pkgs) // -> Tested , and it is working well !
			fmt.Println("This command is just for testing purposes.")
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
