package tui

import (
	"depoty/internal/deletion"
	"depoty/internal/installation"
	"depoty/internal/updation"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TuiStart() {

	// Creation of the application.
	app := tview.NewApplication()

	// Installed Packages Table.
	packageTable := tview.NewTable()
	listOfPkgs := ListWholePkgs(packageTable)

	// the (flex Box,search Bar,searching Text).
	SearchFlex, installPkg, textPkgs := searchPkgs()
	// List found packages Table.
	foundPkgsTable := ListFoundPkgs()

	// Search and packages containers Flex
	pkgsFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(SearchFlex, 0, 2, false).
		AddItem(foundPkgsTable, 0, 5, false).
		AddItem(packageTable, 0, 5, false)

	// Display the information of a package.
	pkgInfo := tview.NewTextView().SetText("Press Enter to fetch a package information.").SetDynamicColors(true)
	pkgInfo.SetBorder(true).SetTitle("ℹ️Package Information")

	// Fetch the installed package information.
	fetchInstalledPkgs(app, pkgInfo, packageTable)

	// Fetch the found package information.
	fetchFoundPkgs(app, pkgInfo, packageTable)

	// Start searching by pressing Enter.
	performSearchingOperation(installPkg, foundPkgsTable)

	// Dynamic package name showing
	dynamicSearchingText(foundPkgsTable, textPkgs)

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
	InstallConfModal := OperationOnPackage("Are you sure you want to install this package?", installation.InstallPkg, app, foundPkgsTable, rowFlex)

	// Handle the installation button & press (To Trigger the process of installation)
	TriggerInstallProcess(tcell.KeyCtrlD, app, foundPkgsTable, InstallConfModal)

	// Create Confirmation box & Handle the updation process
	UpdateconfModal := OperationOnPackage("Are you sure you want to Update this package?", updation.UpdatePkg, app, packageTable, rowFlex)

	// Create Confirmation box & Handle the Deletion process
	DeleteconfModal := OperationOnPackage("Are you sure you want to Delete this package?", deletion.DeletePkg, app, packageTable, rowFlex)

	// Slice of keys
	keysOfOperation := []tcell.Key{tcell.KeyCtrlU, tcell.KeyCtrlQ}

	// Handle the Update & Deletion button & press (To Trigger the process of Updation / Deletion)
	TriggerUpdAndDelProcess(keysOfOperation, app, packageTable, UpdateconfModal, DeleteconfModal)

	// Search in installed packages
	SearchInInstalledPkgs(installPkg, textPkgs, listOfPkgs, packageTable)

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
			app.SetRoot(rowFlex, true).SetFocus(packageTable)
		}

		// Change searching to blank
		if app.GetFocus() != SearchFlex {
			textPkgs.SetText("")
		}

		// Refresh the installed packages
		if event.Key() == tcell.KeyCtrlR {
			ListWholePkgs(packageTable)
		}

		// Already focuesd / no need to edit it
		// var noEditForFound bool = false
		// var noEditForInstalled bool = false

		// // Change the frame text based on focusing components
		// if app.GetFocus() == foundPkgsTable {
		// 	// Show Installation key if it is focused on found packges
		// 	if !noEditForFound {
		// 		guideFrame.Clear().
		// 			AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
		// 			AddText("CTRL + D (Install package)", false, tview.AlignRight, tcell.ColorWhite)
		// 		noEditForFound = true
		// 		noEditForInstalled = false
		// 	}
		// 	// Show the Upgrade and Delete keys if it is focused on Installed Packages
		// } else if app.GetFocus() == packageTable {
		// 	if !noEditForInstalled {
		// 		guideFrame.Clear().
		// 			AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
		// 			AddText("CTRL + U (Update package) / CTRL + Q (Delete package)", false, tview.AlignRight, tcell.ColorWhite)
		// 		noEditForInstalled = true
		// 		noEditForFound = false
		// 	}
		// 	// If the focues otherwise, then clear all and add the basic guide
		// } else {
		// 	guideFrame.Clear().
		// 		AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite)
		// 	noEditForFound = false
		// 	noEditForInstalled = false
		// }

		return event

	})

	// Change the nevigation data if it is focused in packagesTable
	packageTable.SetFocusFunc(func() {
		guideFrame.Clear().
			AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite).
			AddText("CTRL + U (Update package) / CTRL + Q (Delete package)", false, tview.AlignRight, tcell.ColorWhite)
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
