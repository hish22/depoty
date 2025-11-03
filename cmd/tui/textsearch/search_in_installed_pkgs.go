package textsearch

import (
	"strings"

	"github.com/rivo/tview"
)

func SearchInInstalledPkgs(installPkg *tview.InputField, textPkgs *tview.TextView, listOfPkgs map[string]string, packageTable *tview.Table) {
	// Search in installed packages
	installPkg.SetChangedFunc(func(text string) {
		textPkgs.SetText("typing...")
		// var pkgs *map[string]string
		// pkgs = &listOfPkgs
		i := 0
		packageTable.Clear()

		if len(text) != 0 {
			for pkg, totalPkg := range listOfPkgs {
				if strings.Contains(strings.ToLower(totalPkg), strings.ToLower(text)) {
					textPkgs.SetText(pkg)
					packageTable.SetCell(i, 0, tview.NewTableCell(totalPkg))
					i++
					textPkgs.SetText(pkg + " found")
				}
			}
		} else {
			for _, totalPkg := range listOfPkgs {
				packageTable.SetCell(i, 0, tview.NewTableCell(totalPkg))
				i++
			}
		}

	})
}
