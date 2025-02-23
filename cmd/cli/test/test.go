package test

import (
	"depoty/cmd/cli/core"
	"depoty/internal/installation"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "testing",
		Run: func(cmd *cobra.Command, args []string) {
			installation.InstallPkg("php")
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
