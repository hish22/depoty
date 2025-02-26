package messages

import "github.com/rivo/tview"

func MessageModal(app *tview.Application, modal *tview.Modal, root tview.Primitive, message string, title string) *tview.Modal {
	modal.SetText(message).
		SetTitle(title)
	modal.AddButtons([]string{"enter"})

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "enter" {
			app.SetRoot(root, true)
		}
	})

	return modal

}
