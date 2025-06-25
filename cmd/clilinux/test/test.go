//go:build linux

package test

import (
	"depoty/cmd/clilinux/core"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "Testing.",
		Run: func(cmd *cobra.Command, args []string) {
			println("Testing command in linux")
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
