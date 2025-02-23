package tui

import "github.com/rivo/tview"

func fetchInstalledPkgs(app *tview.Application, pkgInfo *tview.TextView, packageTable *tview.Table) {
	// Fetch the installed package information.
	packageTable.SetSelectedFunc(func(row, column int) {
		startFindOper(app, pkgInfo, packageTable)
	})
}
