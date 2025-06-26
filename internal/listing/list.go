package listing

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"runtime"
	"strings"
)

func ListPkgs() []string {

	switch runtime.GOOS {
	case "windows":
		return listchoco()
	case "linux":
		return listapt()
	default:
		return make([]string, 0)
	}

}

func listchoco() []string {
	// Open DB
	db := badgers.MainDb("/tmp/choco/outdate")

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

func listapt() []string {

	// Open DB
	db := badgers.MainDb("/tmp/apt/outdate")

	// Close DB
	defer db.Close()

	// Execute apt list to show installed packages
	text := common.ExecuteScript("apt list --installed", "")
	// new slice splited by \n
	var installed []string = strings.Split(text, "\n")
	// empty slice of strings
	var installed_names []string
	// loop throguh names and split by (/)
	for i, v := range installed {
		name := strings.Split(v, "/")
		installed_names = append(installed_names, name[0])
		_, err := badgers.Read(db, []byte(name[0]))
		// If found outdated package, then add outdated tag.
		if err == nil {
			installed_names[i] += " (outdated)"
		}
	}
	// return the installed packages
	return installed_names[1 : len(installed_names)-1]
}
