package tui

import "github.com/rivo/tview"

func SearchInInstalledPkgs(installPkg *tview.InputField, textPkgs *tview.TextView, listOfPkgs map[string]string, packageTable *tview.Table) {
	// Search in installed packages
	installPkg.SetChangedFunc(func(text string) {
		textPkgs.SetText("Searching...")
		// var pkgs *map[string]string
		// pkgs = &listOfPkgs
		for pkg, totalPkg := range listOfPkgs {
			if text == pkg {
				textPkgs.SetText(pkg)
				packageTable.Clear()
				packageTable.SetCell(0, 0, tview.NewTableCell(totalPkg))
			}
		}

	})
}
