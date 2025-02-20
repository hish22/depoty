package ai

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func SetAiApiCommand() {
	var SetApiCommand = &cobra.Command{
		Use:   "setapi",
		Short: "Set AI agent api key.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty setapi <api_key>")
		},
	}
	core.RootCommand.AddCommand(SetApiCommand)
}
