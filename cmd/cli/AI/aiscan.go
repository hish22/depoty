package ai

import (
	"depoty/cmd/cli/core"
	"fmt"

	"github.com/spf13/cobra"
)

func AiScanCommand() {

	var installWAi bool

	var SetApiCommand = &cobra.Command{
		Use:   "aiscan",
		Short: "find a specific package using AI agent.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("depoty aiscan <package>")
		},
	}

	SetApiCommand.Flags().BoolVarP(&installWAi, "install", "i", false, "enforce direct installation through ai agent.")

	core.RootCommand.AddCommand(SetApiCommand)
}
