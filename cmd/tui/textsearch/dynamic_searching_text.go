package textsearch

import (
	"github.com/rivo/tview"
)

func DynamicSearchingText(foundPkgsTable *tview.Table, textPkgs *tview.TextView) {
	// Dynamic package name showing
	foundPkgsTable.SetSelectionChangedFunc(func(row, column int) {
		textPkgs.SetText(foundPkgsTable.GetCell(row, 0).Text)
	})
}
