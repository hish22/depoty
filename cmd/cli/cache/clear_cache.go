package cache

import (
	"depoty/cmd/cli/core"
	"depoty/internal/badgers"
	"fmt"

	"github.com/spf13/cobra"
)

func ClearCahce() {

	var clearCache = &cobra.Command{
		Use:   "clear",
		Short: "Clear the system cache",
		Run: func(cmd *cobra.Command, args []string) {
			db1 := badgers.MainDb("/tmp/choco/outdate")
			db2 := badgers.MainDb("/system/choco/config")

			fmt.Println("Clearing Cache..")

			db1.DropAll()
			db2.DropAll()

			db1.Close()
			db2.Close()

			fmt.Println("Cache is cleared successfully.")
		},
	}
	core.RootCommand.AddCommand(clearCache)
}
