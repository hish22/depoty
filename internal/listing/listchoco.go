package listing

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"strings"
)

func ListPkgs() []string {

	// Open DB
	db := badgers.MainDb("/tmp/badger/outdate")

	// Close DB
	defer db.Close()

	// perform list "choco list -r" to list packages.
	text := common.ExecuteScript("choco list -r", "")

	// seperate packages by new line.
	listSlice := strings.Split(text, "\n")
	// Final slice.
	listSliceSep := make([]string, 0, 100)

	// checking the list if it has outdated packages.
	for i, item := range listSlice {
		// seperate packages name and version by space: (nmap|1.1.1) -> (nmap 1.1.1).
		pkgName := strings.Split(item, "|")[0]
		// add the seperated value to the slice.
		listSliceSep = append(listSliceSep, strings.Replace(item, "|", " ", 1))
		// Check if there are some outdated packages.
		_, err := badgers.Read(db, []byte(pkgName))
		// If found outdated package, then add outdated tag.
		if err == nil {
			listSliceSep[i] += " (outdated)"
		}

	}

	return listSliceSep
}
