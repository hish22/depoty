package tui

import (
	"depoty/cmd/tui/fetch"
	"depoty/cmd/tui/list"
	"depoty/cmd/tui/operation"
	"depoty/cmd/tui/textsearch"
	"depoty/internal/deletion"
	"depoty/internal/installation"
	"depoty/internal/updation"
	"fmt"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TuiStart() {

	// Creation of the application.
	app := tview.NewApplication()

	// Installed Packages Table.
	packageTable := tview.NewTable()
	listOfPkgs := list.ListWholePkgs(packageTable)

	// the (flex Box,search Bar,searching Text).
	SearchFlex, installPkg, textPkgs := textsearch.SearchPkgs()
	// List found packages Table.
	foundPkgsTable := list.ListFoundPkgs()

	// Search and packages containers Flex
	pkgsFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(SearchFlex, 0, 2, false).
		AddItem(foundPkgsTable, 0, 5, false).
		AddItem(packageTable, 0, 5, false)

	// Display the information of a package.
	pkgInfo := tview.NewTextView().SetText("Press Enter to fetch a package information.").SetDynamicColors(true)
	pkgInfo.SetBorder(true).SetTitle("ℹ️Package Information")

	// Fetch the installed package information.
	fetch.FetchInstalledPkgs(app, pkgInfo, packageTable)

	// Fetch the found package information.
	fetch.FetchFoundPkgs(app, pkgInfo, foundPkgsTable)

	// Start searching by pressing Enter.
	operation.PerformSearchingOperation(installPkg, foundPkgsTable)

	// Dynamic package name showing
	textsearch.DynamicSearchingText(foundPkgsTable, textPkgs)

	// Create Modal and page to Show Keys information
	keyInfo := guideInfoModalPage()

	// Guide Info components bottom bar / First Iteration
	guideFrame := tview.NewFrame(tview.NewBox()).
		AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
		AddText("CTRL + U (Update package) / CTRL + Q (Delete package)", false, tview.AlignRight, tcell.ColorWhite)

	// Create the flex container, to contain the left side and right side
	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(pkgsFlex, 0, 1, false).
		AddItem(pkgInfo, 0, 2, false)
	// Main Flex , the total flex
	rowFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 15, false).
		AddItem(guideFrame, 0, 2, false)

	// Create Confirmation box & Handle the installation process
	InstallConfModal := operation.OperationOnPackage("Are you sure you want to install this package?", installation.InstallPkg, app, foundPkgsTable, rowFlex, "Package successfully installed", "Package installation failed")

	// Handle the installation button & press (To Trigger the process of installation)
	TriggerInstallProcess(tcell.KeyCtrlD, app, foundPkgsTable, InstallConfModal)

	// Create Confirmation box & Handle the updation process
	UpdateconfModal := operation.OperationOnPackage("Are you sure you want to Update this package?", updation.UpdatePkg, app, packageTable, rowFlex, "Package successfully updated", "Package update failed")

	// Create Confirmation box & Handle the Deletion process
	DeleteconfModal := operation.OperationOnPackage("Are you sure you want to Delete this package?", deletion.DeletePkg, app, packageTable, rowFlex, "Package successfully deleted", "Package deletion failed")

	// Create Confirmation box & Handle the drop packages process
	var DropconfModal *tview.Modal
	if runtime.GOOS == "windows" {
		DropconfModal = operation.DropAllPkgsOperation("Are you sure you want to Delete all installed packges?", deletion.DropAllPkgs, app, packageTable, rowFlex, "Packages successfully dropped", "Failed to drop packages")
	}

	// Create Confirmation box & Handle the Update all packages process
	UpgradeconfModal := operation.UpdateAllPkgsOperation("Are you sure you want to Update all installed packages?", updation.UpdateAllPkgs, app, packageTable, rowFlex, "Packages successfully updated", "Failed to Update packages")

	// Specifiy update all key based on platform
	var updateKey tcell.Key
	switch runtime.GOOS {
	case "windows":
		updateKey = tcell.KeyF10
	case "linux":
		updateKey = tcell.KeyF2
	}

	// Slice of keys
	keysOfOperation := []tcell.Key{tcell.KeyCtrlU, tcell.KeyCtrlQ, tcell.KeyF12, updateKey}

	// Handle the Update & Deletion button & press (To Trigger the process of Updation / Deletion)
	TriggerUpdAndDelProcess(keysOfOperation, app, packageTable, UpdateconfModal, DeleteconfModal, DropconfModal, UpgradeconfModal)

	// Nabigation and focues change + Handle refresh trigger
	// Also change the text under the search to blank if searchFlex losed focus
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		// Specific Navigation keys
		if event.Key() == tcell.KeyCtrlS {
			app.SetFocus(installPkg)
		} else if event.Key() == tcell.KeyCtrlP {
			app.SetFocus(packageTable)
		} else if event.Key() == tcell.KeyCtrlN {
			app.SetFocus(pkgInfo)
		} else if event.Key() == tcell.KeyCtrlF {
			app.SetFocus(foundPkgsTable)
		}

		// General navigation
		if app.GetFocus() == packageTable {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(foundPkgsTable)
			}
		} else if app.GetFocus() == foundPkgsTable {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(SearchFlex)
			}
		} else if app.GetFocus() == SearchFlex {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(pkgInfo)
			}
		} else if app.GetFocus() == pkgInfo {
			if event.Key() == tcell.KeyTab {
				app.SetFocus(packageTable)
			}
		}

		// Show Key info
		if event.Key() == tcell.KeyF1 {
			app.SetRoot(keyInfo, true).SetFocus(keyInfo)
		}

		// Set the focus to the program
		if event.Key() == tcell.KeyEsc {
			// get the app main focus
			focus := app.GetFocus()

			// Make the focues to the whole search flex / the package table if focus was on guide frame
			if focus == installPkg {
				focus = SearchFlex
			} else {
				focus = packageTable
			}
			// Reset the app root and focus to the main row flex
			app.SetRoot(rowFlex, true).SetFocus(focus)
		}

		// Change searching to blank
		if app.GetFocus() != SearchFlex {
			textPkgs.SetText("")
		}

		// Refresh the installed packages
		if event.Key() == tcell.KeyCtrlR {
			packageTable.Clear()
			Contentlist := list.RefreshWholePkgs(packageTable, &listOfPkgs)
			for i, content := range Contentlist {
				packageTable.SetCell(i, 0, tview.NewTableCell(content))
			}
			app.Suspend(func() {
				fmt.Println("Refreshing App...")
			})
		}

		return event

	})

	// Search in installed packages
	textsearch.SearchInInstalledPkgs(installPkg, textPkgs, listOfPkgs, packageTable)

	// Change the nevigation data if it is focused in packagesTable
	packageTable.SetFocusFunc(func() {
		if runtime.GOOS == "windows" {
			guideFrame.Clear().
				AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
				AddText("CTRL + U (Update) / CTRL + Q (Delete) / F9 (Drop) / F10 (upgrade)", false, tview.AlignRight, tcell.ColorWhite)
		} else {
			guideFrame.Clear().
				AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
				AddText("CTRL + U (Update) / CTRL + Q (Delete) / F2 (upgrade)", false, tview.AlignRight, tcell.ColorWhite)
		}
	})
	// Change the nevigation data if it is focused in foundPkgsTable
	foundPkgsTable.SetFocusFunc(func() {
		guideFrame.Clear().
			AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
			AddText("CTRL + D (Install package)", false, tview.AlignRight, tcell.ColorWhite)
	})
	// Change the nevigation data if it is focused in SearchFlex (Default)
	SearchFlex.SetFocusFunc(func() {
		guideFrame.Clear().
			AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite)
	})
	// Change the nevigation data if it is focused in pkgInfo (Default)
	pkgInfo.SetFocusFunc(func() {
		guideFrame.Clear().
			AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite)
	})

	// Set focus to the dropdown instead of the flex / Start the event loop
	if err := app.SetRoot(rowFlex, true).SetFocus(packageTable).Run(); err != nil {
		panic(err)
	}

}
