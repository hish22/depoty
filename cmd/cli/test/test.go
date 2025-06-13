package test

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "Testing.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is testing command!")
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
