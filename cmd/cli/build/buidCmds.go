package build

import (
	ai "depoty/cmd/cli/AI"
	"depoty/cmd/cli/category"
	"depoty/cmd/cli/core"
	"depoty/cmd/cli/decomand"
	"depoty/cmd/cli/initd"
	"depoty/cmd/cli/install"
	"depoty/cmd/cli/list"
	"depoty/cmd/cli/revert"
	"depoty/cmd/cli/search"
	"depoty/cmd/cli/update"
	"log"
	"os"
)

func CreateCommands() {

	// Init the "init" command.
	initd.MainInitCommand()

	// Init the "install" command
	install.InstallCommand()

	// Init the "update" command
	update.UpdateCommand()

	// Init the "delete" command
	decomand.DeleteCommand()

	// Init the "revert" command
	revert.RevertCommand()

	// Init the "list" command
	list.ListCommand()

	// Init the "search" command
	search.SearchCommand()

	//Init catagory commands
	// Create cate command
	category.CreateCateCommand()
	// Drop cate command
	category.DropCateCommand()
	// Remove package from cate command
	category.RemovePkgInCateCommand()
	// Find specific package in cate command
	category.FingPkgInCateCommand()

	// AI commands
	// set Api key
	ai.SetAiApiCommand()
	// scan and find package
	ai.AiScanCommand()

	// Executing the main Command
	// Error check regarding the root command
	if err := core.RootCommand.Execute(); err != nil {
		log.Fatal("unpredictable manner happend!")
		os.Exit(1)
	}

}
