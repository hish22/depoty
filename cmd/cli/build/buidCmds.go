package build

import (
	"depoty/cmd/cli/core"
	"depoty/cmd/cli/initd"
	"depoty/cmd/cli/test"
	"log"
	"os"
)

func CreateCommands() {

	// Init the "init" command.
	initd.MainInitCommand()

	//Test commands
	test.TestCommand()

	// Executing the main Command
	// Error check regarding the root command
	if err := core.RootCommand.Execute(); err != nil {
		log.Fatal("unpredictable manner happend!")
		os.Exit(1)
	}

}
