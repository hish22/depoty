package listing

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"strings"
)

func OutdatedList() {
	// open Badger
	db := badgers.MainDb("/tmp/badger/outdate")

	// Close Badger
	defer db.Close()

	// List Outdated packges
	outdatedPackages := common.ExecuteScript("choco outdated -r", "")

	// Split by Lines
	outdatedByLine := strings.Split(outdatedPackages, "\n")

	// Split by (|) , then save to badger
	for i := 0; i < len(outdatedByLine)-1; i++ {
		pkgName := strings.Split(outdatedByLine[i], "|")[0]
		// Check If the value doesn't exsist.
		value, err := badgers.Read(db, []byte(pkgName))
		// Add the package to the db if it is outdated.
		if err != nil {
			item := []byte(pkgName)
			badgers.Insert(db, item, value)
		}

	}

}
