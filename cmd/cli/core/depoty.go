package core

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "depoty",
	Short: "Root command for managing Depoty.",
	Long:  "Depoty is an advanced extension of package managers, offering enhanced local management features for package handling and control.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(`Depoty is an advanced extension of package managers, offering enhanced local management features for package handling and control.
		
	default usage:
		depoty <command> -opitions

	Main Commands:
		depoty  "Root command for managing Depoty."
		depoty init "Initializes Depoty and sets up the package management environment."
		depoty install <Packages>  "Installs specified packages."
		depoty update <Packages> "Updates specified packages to the latest version."
		depoty delete <Packages>  "Removes specified packages."
		depoty revert <Package> <version> "Rolls back a package to a specified version."
		depoty list "Displays all monitored packages."
		depoty search <package>  "Find a specific package"
		
	Pkgs cate Order:	
		depoty create <catagory name> "Create a catagory."
		depoty remove <package> <category> "Remove a specific package from a catagory."
		depoty drop <catagory> "Deletes the entire catagory."
		depoty find <package> <category> "Check a package in specific catagory"

	Ai agent commands:
		depoty setapi <api_key> "Set AI agent api key."
		depoty aiscan <package> "find a specific package using AI agent."
		
		
	Ai flages:
	-i "enforce direct installation of the package."	

	Flags:
		-h "Retrieve Depoty Manual."
		-winget "enforce winget."
		-choco "enforce chocolatey."
		-aiagent "enforce an Ai agent."
		-all "enforce all operation"
		`)
	},
}
