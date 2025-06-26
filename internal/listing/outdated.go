package listing

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"runtime"
	"strings"
)

func OutdatedList() {
	switch runtime.GOOS {
	case "windows":
		outdatechoco()
	case "linux":
		Outdatedapt()
	}
}

func outdatechoco() {
	// open Badger
	db := badgers.MainDb("/tmp/choco/outdate")

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

func Outdatedapt() {

	// open Badger
	db := badgers.MainDb("/tmp/apt/outdate")

	// Close Badger
	defer db.Close()

	// Execute apt list to show outdated packages
	text := common.ExecuteScript("apt list --upgradable", "")
	// new slice splited by \n
	var outdated []string = strings.Split(text, "\n")
	// empty slice of strings
	// var outdated_names []string
	// loop throguh names and split by (/)
	var i int
	for i = 1; i < len(outdated)-1; i++ {
		name := strings.Split(outdated[i], "/")[0]
		// outdated_names = append(outdated_names, name[0])
		// Check If the value doesn't exsist.
		value, err := badgers.Read(db, []byte(name))
		if err != nil {
			item := []byte(name)
			badgers.Insert(db, item, value)
		}
	}

}
