package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TriggerInstallProcess(key tcell.Key, app *tview.Application, PkgsTable *tview.Table, InstallConfModal *tview.Modal) {
	// Handle the Operation button & press
	PkgsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Start the installation process if CTRL + D pressed
		if event.Key() == key {
			// No installation if found packages is empty
			if PkgsTable.GetRowCount() != 0 {
				// Show Confirmation box
				app.SetRoot(InstallConfModal, true).SetFocus(InstallConfModal)
			}
		}
		return event
	})
}

func TriggerUpdAndDelProcess(key []tcell.Key, app *tview.Application, PkgsTable *tview.Table,
	UpdateConfModal *tview.Modal, DeleteModal *tview.Modal, DropAllModal *tview.Modal) {
	// Handle the Operation button & press
	PkgsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Start the Update process if CTRL + U pressed
		if event.Key() == key[0] {
			// No Updation if installed packages is empty
			if PkgsTable.GetRowCount() != 0 {
				// Show Confirmation box
				app.SetRoot(UpdateConfModal, true).SetFocus(UpdateConfModal)
			}
			// Start the Delete process if CTRL + Q pressed
		} else if event.Key() == key[1] {
			// No Deletion if installed packages is empty
			if PkgsTable.GetRowCount() != 0 {
				// Show Confirmation box
				app.SetRoot(DeleteModal, true).SetFocus(DeleteModal)
			}
			// Start the Drop process if F9 pressed
		} else if event.Key() == key[2] {
			// No Droping if installed packages is empty
			if PkgsTable.GetRowCount() != 0 {
				// Show Confirmation box
				app.SetRoot(DropAllModal, true).SetFocus(DropAllModal)
			}
		}
		return event
	})
}
