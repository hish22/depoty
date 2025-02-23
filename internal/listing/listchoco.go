package listing

import (
	"depoty/internal/badgers"
	"depoty/internal/util/common"
	"regexp"
	"strings"
)

func ListPkgs() []string {

	// Open DB
	db := badgers.MainDb("/tmp/badger/outdate")

	// Close DB
	defer db.Close()

	// List All packages
	text := common.ExecuteScript("choco list", "")

	// Slice the listed packages
	var textSlice []string = make([]string, 0, 100)

	// Match the regular expression
	re := regexp.MustCompile(`(\S+)\s+([\d.]+)`)
	match := re.FindAllStringSubmatch(text, -1)

	// Passing the list to textSlice
	for _, item := range match {
		textSlice = append(textSlice, item...)
	}
	// checking the list if it has outdated packages
	for i, item := range textSlice {
		pkgName := strings.Split(item, " ")[0]

		_, err := badgers.Read(db, []byte(pkgName))

		if err == nil {
			textSlice[i] += " (outdated)"
		}

	}

	return textSlice
}
