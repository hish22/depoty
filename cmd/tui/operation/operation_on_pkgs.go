package operation

import (
	"depoty/cmd/tui/messages"
	"strings"

	"github.com/rivo/tview"
)

func OperationOnPackage(msg string, operationCallback func(string) bool, app *tview.Application, PkgsTable *tview.Table, rowFlex *tview.Flex, sucessMsg string, errMsg string) *tview.Modal {
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
				var pkgRowSlice []string = strings.Split(pkgRow, " ")
				// Operate on package
				// sucess := callback(pkgRowSlice[0]) // installation.InstallPkg(pkgRowSlice[0])
				var success bool
				app.Suspend(func() {
					success = operationCallback(pkgRowSlice[0]) // ex: installation.InstallPkg(pkgRowSlice[0])
				})
				if success {
					// Update the Installed Packages list
					// list.ListWholePkgs(PkgsTable)
					// Update the view
					// Success modal message.
					successModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, sucessMsg, "Success message")
					app.SetRoot(successModal, true).SetFocus(successModal)
				} else {
					// Error modal message.
					errorModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, errMsg, "Error message")
					app.SetRoot(errorModal, true).SetFocus(errorModal)
				}

			} else {
				app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
			}
		})
}

func UpdateAllPkgsOperation(msg string, operationCallback func() bool, app *tview.Application, PkgsTable *tview.Table, rowFlex *tview.Flex, sucessMsg string, errMsg string) *tview.Modal {
	// Create Confimation box & Handle the installation process
	return tview.NewModal().
		SetText(msg). // "Are you sure you want to install this package?"
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				// Create Packges slice
				// var PkgSlice []string
				// Start inserting packages into the slice to be updated,
				// for i := 0; i < PkgsTable.GetRowCount(); i++ {
				// 	pkgRow := PkgsTable.GetCell(i, 0)
				// 	pkgText := pkgRow.Text
				// 	if strings.HasSuffix(pkgText, "(outdated)") {
				// 		splitBySpace := strings.Split(pkgRow.Text, " ")
				// 		PkgSlice = append(PkgSlice, splitBySpace[0])

				// 	}
				// }
				// if len(PkgSlice) == 0 {
				// 	allUpdatedModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, "All packages are updated", "All Updated message")
				// 	app.SetRoot(allUpdatedModal, true).SetFocus(allUpdatedModal)
				// }

				var success bool
				app.Suspend(func() {
					//success = operationCallback(PkgSlice) // ex: installation.InstallPkg(pkgRowSlice[0])
					success = operationCallback()
				})
				if success {
					// Update the Installed Packages list
					// list.ListWholePkgs(PkgsTable)
					// Update the view
					// Success modal message.
					successModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, sucessMsg, "Success message")
					app.SetRoot(successModal, true).SetFocus(successModal)
				} else {
					// Error modal message.
					errorModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, errMsg, "Error message")
					app.SetRoot(errorModal, true).SetFocus(errorModal)
				}

			} else {
				app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
			}
		})
}

func DropAllPkgsOperation(msg string, operationCallback func([]string) bool, app *tview.Application, PkgsTable *tview.Table, rowFlex *tview.Flex, sucessMsg string, errMsg string) *tview.Modal {
	// Create Confimation box & Handle the installation process
	return tview.NewModal().
		SetText(msg). // "ex: Are you sure you want to install this package?"
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				// Create Packges slice
				var PkgSlice []string
				// Start inserting packages into the slice to be deleted,
				// If the package contain chocolatey or depoty, then don't delete it.
				for i := 0; i < PkgsTable.GetRowCount(); i++ {
					pkgRow := PkgsTable.GetCell(i, 0)
					splitBySpace := strings.Split(pkgRow.Text, " ")
					// If package is not chocolatey, then continue deleting
					if strings.ToLower(splitBySpace[0]) == "chocolatey" {
						continue
					}
					// If package is not depoty, then continue deleting
					if strings.ToLower(splitBySpace[0]) == "depoty" {
						continue
					}
					// If it is not a chocolatey extension, then continue deleting
					if strings.Contains(strings.ToLower(splitBySpace[0]), "chocolatey") {
						continue
					}
					// Delete packages
					PkgSlice = append(PkgSlice, splitBySpace[0])
				}
				// Drop All package
				var success bool
				app.Suspend(func() {
					success = operationCallback(PkgSlice) // ex: deletion.DropAllPkgs(pkgRowSlice)
				})

				if success {
					// Update the Installed Packages list
					// ListWholePkgs(packageTable)
					// Update the view
					// Success modal message.
					successModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, sucessMsg, "Success message")
					app.SetRoot(successModal, true).SetFocus(successModal)
				} else {
					// Error modal message.
					errorModal := messages.MessageModal(app, tview.NewModal(), rowFlex, PkgsTable, errMsg, "Error message")
					app.SetRoot(errorModal, true).SetFocus(errorModal)
				}

			} else {
				app.SetRoot(rowFlex, true).SetFocus(PkgsTable)
			}
		})
}
