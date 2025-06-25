//go:build windows

package initd

import (
	"depoty/cmd/cli/core"
	"depoty/internal/initalization"

	"github.com/spf13/cobra"
)

func MainInitCommand() {
	var InitCommand = &cobra.Command{
		Use:   "init",
		Short: "Initializes Depoty and sets up the package management environment.",
		Run: func(cmd *cobra.Command, args []string) {
			initalization.EntryPoint()
		},
	}
	core.RootCommand.AddCommand(InitCommand)
}
