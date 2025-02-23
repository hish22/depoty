package tui

import (
	"strings"

	"github.com/rivo/tview"
)

func OperationOnPackage(msg string, callback func(string) bool, app *tview.Application, PkgsTable *tview.Table, rowFlex *tview.Flex) *tview.Modal {
	// Create Confimation box & Handle the installation process
	return tview.NewModal().
		SetText(msg). // "Are you sure you want to install this package?"
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				// Get package name
				r, _ := PkgsTable.GetSelection()
				// Find the selected row
				pkgRow := PkgsTable.GetCell(r, 0).Text
				// Split the text into slices to get the package name
				pkgRowSlice := strings.Split(pkgRow, " ")
				// Operate on package
				// sucess := callback(pkgRowSlice[0]) // installation.InstallPkg(pkgRowSlice[0])
				var success bool
				app.Suspend(func() {
					success = callback(pkgRowSlice[0]) // ex: installation.InstallPkg(pkgRowSlice[0])
				})
				if success {
					// Update the Installed Packages list
					// ListWholePkgs(packageTable)
					// Update the view
					app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
				} else {
					app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
				}

			} else {
				app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
			}
		})
}

func DropAllPkgsOperation(msg string, callback func([]string) bool, app *tview.Application, PkgsTable *tview.Table, rowFlex *tview.Flex) *tview.Modal {
	// Create Confimation box & Handle the installation process
	return tview.NewModal().
		SetText(msg). // "ex: Are you sure you want to install this package?"
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				// Create Packges slice
				var PkgSlice []string

				for i := 0; i < PkgsTable.GetRowCount(); i++ {
					pkgRow := PkgsTable.GetCell(i, 0)
					splitBySpace := strings.Split(pkgRow.Text, " ")
					if !strings.Contains(splitBySpace[0], "chocolatey") {
						PkgSlice = append(PkgSlice, splitBySpace[0])
					}
				}
				// Drop All package
				var success bool
				app.Suspend(func() {
					success = callback(PkgSlice) // ex: installation.InstallPkg(pkgRowSlice[0])
				})

				if success {
					// Update the Installed Packages list
					// ListWholePkgs(packageTable)
					// Update the view
					app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
				} else {
					app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
				}

			} else {
				app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
			}
		})
}
