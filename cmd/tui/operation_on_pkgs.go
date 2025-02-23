package tui

import (
	"strings"

	"github.com/rivo/tview"
)

func OperationOnPackage(msg string, callback func(string) bool, app *tview.Application, foundPkgsTable *tview.Table, packageTable *tview.Table, rowFlex *tview.Flex) *tview.Modal {
	// Create Confimation box & Handle the installation process
	return tview.NewModal().
		SetText(msg). // "Are you sure you want to install this package?"
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				// Get package name
				r, _ := foundPkgsTable.GetSelection()
				// Find the selected row
				pkgRow := foundPkgsTable.GetCell(r, 0).Text
				// Split the text into slices to get the package name
				pkgRowSlice := strings.Split(pkgRow, " ")
				// Install package
				sucess := callback(pkgRowSlice[0]) // installation.InstallPkg(pkgRowSlice[0])

				if sucess {
					// Update the Installed Packages list
					ListWholePkgs(packageTable)
					// Update the view
					app.SetRoot(rowFlex, true).SetFocus(foundPkgsTable)
				} else {
					app.SetRoot(rowFlex, true).SetFocus(foundPkgsTable)
				}

			} else {
				app.SetRoot(rowFlex, true).SetFocus(foundPkgsTable)
			}
		})
}
