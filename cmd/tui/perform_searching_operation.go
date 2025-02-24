package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func performSearchingOperation(app *tview.Application, installPkg *tview.InputField, foundPkgsTable *tview.Table, textPkg *tview.TextView) {
	// Start searching by pressing Enter
	installPkg.SetDoneFunc(func(key tcell.Key) {

		if installPkg.GetText() == "" {
			return
		}

		operateFindingPkgs(installPkg, foundPkgsTable)
	})
}
