//go:build windows

package test

import (
	"depoty/cmd/cli/core"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "Testing.",
		Run: func(cmd *cobra.Command, args []string) {
			println("This is test command!")
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
