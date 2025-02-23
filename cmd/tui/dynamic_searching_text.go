package tui

import "github.com/rivo/tview"

func dynamicSearchingText(foundPkgsTable *tview.Table, textPkgs *tview.TextView) {
	// Dynamic package name showing
	foundPkgsTable.SetSelectionChangedFunc(func(row, column int) {
		textPkgs.SetText(foundPkgsTable.GetCell(row, 0).Text)
	})
}
