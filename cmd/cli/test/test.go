package test

import (
	"depoty/cmd/cli/core"
	"depoty/internal/finding"
	"fmt"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "testing",
		Run: func(cmd *cobra.Command, args []string) {
			text := finding.FindPkg("dsdsds")

			fmt.Println(text)
			fmt.Println(len(text))

		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
