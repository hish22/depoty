//go:build linux

package test

import (
	"depoty/cmd/clilinux/core"
	"depoty/internal/util/common"
	"strings"

	"github.com/spf13/cobra"
)

func TestCommand() {
	var TestCommand = &cobra.Command{
		Use:   "test",
		Short: "Testing.",
		Run: func(cmd *cobra.Command, args []string) {
			text := common.ExecuteScript("apt list --upgradable", "")
			var installed []string = strings.Split(text, "\n")
			var installed_names []string
			for _, v := range installed {
				name := strings.Split(v, "/")
				installed_names = append(installed_names, name[0])
			}
			for _, v := range installed_names[0 : len(installed_names)-1] {
				println(v)
			}
			// println("Testing command in linux")
		},
	}
	core.RootCommand.AddCommand(TestCommand)
}
