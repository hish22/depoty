package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func performSearchingOperation(installPkg *tview.InputField, foundPkgsTable *tview.Table) {
	// Start searching by pressing Enter
	installPkg.SetDoneFunc(func(key tcell.Key) {

		if installPkg.GetText() == "" {
			return
		}

		operateFindingPkgs(installPkg, foundPkgsTable)
	})
}
