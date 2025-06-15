package fetch

import "github.com/rivo/tview"

func FetchInstalledPkgs(app *tview.Application, pkgInfo *tview.TextView, packageTable *tview.Table) {
	// Fetch the installed package information.
	packageTable.SetSelectedFunc(func(row, column int) {
		startFindOper(app, pkgInfo, packageTable)
	})
}
