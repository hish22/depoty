package install

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func InstallCommand() {
	var InstallCommand = &cobra.Command{
		Use:   "install",
		Short: "Installs specified packages.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty install <Packages>")
		},
	}
	core.RootCommand.AddCommand(InstallCommand)
}
