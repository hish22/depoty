package fetch

import "github.com/rivo/tview"

func FetchFoundPkgs(app *tview.Application, pkgInfo *tview.TextView, foundPkgsTable *tview.Table) {
	// Fetch the found package information.
	foundPkgsTable.SetSelectedFunc(func(row, column int) {
		startFindOper(app, pkgInfo, foundPkgsTable)
	})
}
