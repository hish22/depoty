package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TriggerOperationProcess(key tcell.Key, app *tview.Application, PkgsTable *tview.Table, confModal *tview.Modal) {
	// Handle the Operation button & press
	PkgsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Start the installation process if CTRL + D pressed
		if event.Key() == key {
			// No installation if found packages is empty
			if PkgsTable.GetRowCount() != 0 {
				// Show Confirmation box
				app.SetRoot(confModal, true).SetFocus(confModal)
			} else {

			}
		}
		return event
	})
}
