package listing

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"fmt"
	"strings"
)

func OutdatedList() {
	// open Badger
	db := badgers.MainDb()

	// Close Badger
	defer db.Close()

	// List Outdated packges
	outdatedPackages := common.ExecuteScript("choco outdated", "")

	// Split by Lines
	outdatedByLine := strings.Split(outdatedPackages, "\n")

	// Split by (|) , then save to badger
	for i := 4; i < len(outdatedByLine); i++ {
		splitVersionAndName := strings.Split(outdatedByLine[i], "|")

		if i != len(outdatedByLine)-2 && splitVersionAndName[0] != "" {
			value, err := badgers.Read(db, []byte(splitVersionAndName[0]))
			fmt.Println(string(value))
			if err != nil {
				item := []byte(splitVersionAndName[0])
				badgers.Insert(db, item, item)
			}
		}

	}

}
