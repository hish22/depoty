package tui

import (
	"depoty/internal/listing"
	"strings"

	"github.com/rivo/tview"
)

func ListWholePkgs(packageTable *tview.Table) map[string]string {

	list := make(map[string]string)

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := listing.ListPkgs()
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	// Row Counter
	j := 0

	// Packges Slice
	// Create the packageTable
	for i := 0; i < len(pkgs); i += 3 {
		packageTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
		list[strings.Split(pkgs[i], " ")[0]] = pkgs[i]
		j++
	}

	packageTable.SetBorder(true)
	packageTable.SetTitle("📦Installed Packages")

	packageTable.Select(0, 0).
		SetSelectable(true, true)

	return list
}
