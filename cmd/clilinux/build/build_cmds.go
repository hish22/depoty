//go:build linux

package build

import (
	"depoty/cmd/clilinux/core"
	"depoty/cmd/clilinux/test"
	"log"
	"os"
)

func CreateCommands() {

	// Init the "init" command.
	// initd.MainInitCommand()

	// Test commands
	test.TestCommand()

	// Clear Cache command
	// cache.ClearCahce()

	// Executing the main Command
	// Error check regarding the root command
	if err := core.RootCommand.Execute(); err != nil {
		log.Fatal("unpredictable manner happend!")
		os.Exit(1)
	}

}
