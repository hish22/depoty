package tui

import (
	"depoty/internal/finding"
	"depoty/internal/installation"
	"depoty/internal/listing"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ListWholePkgs(packageTable *tview.Table) {

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := listing.ListPkgs()
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	// Row Counter
	j := 0

	// Create the packageTable
	for i := 0; i < len(pkgs); i += 3 {
		packageTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
		j++
	}

	packageTable.SetBorder(true)
	packageTable.SetTitle("📦Installed Packages")

	packageTable.Select(0, 0).
		SetSelectable(true, true)

	// return packageTable
}

func searchPkgs() (*tview.Flex, *tview.InputField, *tview.TextView) {
	//Search bar components to install package
	searchPkg := tview.NewInputField()

	searchPkg.SetLabel("Find Package:")

	searchPkg.SetFieldBackgroundColor(tcell.ColorYellow)

	searchPkg.SetFieldTextColor(tcell.ColorBlack)

	// searchPkg.SetPlaceholder("Package name").
	// 	SetPlaceholderTextColor(tcell.ColorBlack).

	searchPkg.Autocomplete()

	// Searching text
	searchText := tview.NewTextView()

	flexBox := tview.NewFlex().
		AddItem(searchPkg, 0, 1, false).
		AddItem(searchText, 0, 1, false)

	flexBox.
		SetDirection(tview.FlexRow).
		SetBorder(true).
		SetTitle("🔍Search")

	return flexBox, searchPkg, searchText

}

func ListFoundPkgs() *tview.Table {

	packageTable := tview.NewTable()

	packageTable.SetBorder(true)
	packageTable.SetTitle("📥Found Packages")

	packageTable.Select(0, 0).
		SetSelectable(true, true)

	return packageTable

}

func startFindOper(app *tview.Application, pkgInfo *tview.TextView, packageTable *tview.Table) {
	// Clearing text buffer
	pkgInfo.Clear()

	// Print starting fetching
	pkgInfo.SetText("Fetching information...")

	// Force Drawing to refresh the secreen immediately
	app.ForceDraw()

	// Get selected Row/Columns
	r, c := packageTable.GetSelection()

	// Seperate name from version
	textSlice := strings.Split(packageTable.GetCell(r, c).Text, " ")

	// Start Fetching app/service information
	info := finding.FindPkgInfo(textSlice[0])

	infoSplit := strings.Split(info, "\n")

	var resultInfo strings.Builder

	for i, info := range infoSplit {
		if i < 17 && i >= 3 {
			resultInfo.WriteString("\n" + "[yellow]" + info + "[-]\n\n")
		} else {
			resultInfo.WriteString(info + "\n")
		}
	}

	pkgInfo.SetText(resultInfo.String())

}

func operateFoundPkgs(installPkg *tview.InputField, foundPkgsTable *tview.Table) {
	// Clear the found packages content
	foundPkgsTable.Clear()

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := finding.FindPkg(installPkg.GetText())
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	// pkgs := finding.FindPkg(installPkg.GetText())

	j := 0

	for i := 2; i < len(pkgs); i++ {

		if len(pkgs) == 4 || len(pkgs) == 8 {
			if strings.Contains(pkgs[2], "0 packages found.") {
				foundPkgsTable.SetBorderColor(tcell.ColorRed)
				foundPkgsTable.SetTitle(pkgs[2])
				break
			}
		}

		// Get rid of unwanted choco ads (Check this)
		if strings.Contains(pkgs[len(pkgs)-3], "Learn more") && i == len(pkgs)-3 {
			foundPkgsTable.SetTitle(pkgs[i-3])
			foundPkgsTable.SetBorderColor(tcell.ColorGreen)
			break
		} else {
			if len(pkgs)-2 == i {
				foundPkgsTable.SetTitle(pkgs[i])
				foundPkgsTable.SetBorderColor(tcell.ColorGreen)
				break
			}
		}

		foundPkgsTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
		j++
	}
}

func TuiStart() {
	// box := tview.NewBox().SetBorder(true).SetTitle("Depoty")
	app := tview.NewApplication()

	// Installed Packages Table
	packageTable := tview.NewTable()
	ListWholePkgs(packageTable)

	// the (flex Box,search Bar,searching Text)
	SearchFlex, installPkg, textPkgs := searchPkgs()
	// List found packages Table
	foundPkgsTable := ListFoundPkgs()

	// Search and packages Flex
	pkgsFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(SearchFlex, 0, 2, false).
		AddItem(foundPkgsTable, 0, 5, false).
		AddItem(packageTable, 0, 5, false)

	// Display the information of a package
	pkgInfo := tview.NewTextView().SetText("Press Enter to fetch a package information.").SetDynamicColors(true)

	pkgInfo.SetBorder(true).SetTitle("ℹ️Package Information")

	// fetch the installed package information
	packageTable.SetSelectedFunc(func(row, column int) {
		startFindOper(app, pkgInfo, packageTable)
	})

	// Fetch the found package information.
	foundPkgsTable.SetSelectedFunc(func(row, column int) {
		startFindOper(app, pkgInfo, foundPkgsTable)
	})

	// Start searching by pressing Enter
	installPkg.SetDoneFunc(func(key tcell.Key) {

		if installPkg.GetText() == "" {
			return
		}

		operateFoundPkgs(installPkg, foundPkgsTable)
	})

	// Dynamic package name showing
	foundPkgsTable.SetSelectionChangedFunc(func(row, column int) {
		textPkgs.SetText(foundPkgsTable.GetCell(row, 0).Text)
	})

	KeyDataModal := tview.NewModal().SetText(`
		CTRL + S (for search)

		CTRL + F (Found Packages)

		CTRL + P (Installed Packges)

		CTRL + N (Package Info)

		CTRL + R (for Refresh Installed packages)

		Esc (Back)
	`)

	KeyDataModal.SetTitle("Keys Information")

	// Show Keys information
	keyInfo := tview.NewPages().AddPage("Keys Information", KeyDataModal, false, true)

	// Guide Info components
	guideFrame := tview.NewFrame(tview.NewBox()).
		AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite)

	// Create the flex container
	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn). // Explicitly set direction (optional, defaults to Column)
		AddItem(pkgsFlex, 0, 1, false).
		AddItem(pkgInfo, 0, 2, false)
	// Main Flex
	rowFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 15, false).
		AddItem(guideFrame, 0, 2, false)

	// Create Confimation box / Handle the installation process
	confModal := tview.NewModal().
		SetText("Are you sure you want to install this package?").
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
				sucess := installation.InstallPkg(pkgRowSlice[0])

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

	// Handle the installation button & press
	foundPkgsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Start the installation process if CTRL + D pressed
		if event.Key() == tcell.KeyCtrlD {
			// No installation if found packages is empty
			if foundPkgsTable.GetRowCount() != 0 {
				// Get the current Selection
				// r, _ := foundPkgsTable.GetSelection()
				// Show Confirmation box
				app.SetRoot(confModal, true).SetFocus(confModal)
			} else {

			}
		}
		return event
	})

	// Change the components focus
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

		// Refresh the installed packages
		if event.Key() == tcell.KeyCtrlR {
			ListWholePkgs(packageTable)
		}

		// Show Key info
		if event.Key() == tcell.KeyF1 {
			app.SetRoot(keyInfo, true).SetFocus(keyInfo)
		}

		// Set the focus to the program
		if event.Key() == tcell.KeyEsc {
			app.SetRoot(rowFlex, true).SetFocus(packageTable)
		}

		// Already focuesd / no need to edit
		var noEdit bool = false

		// Change the frame text based on focusing components
		if app.GetFocus() == foundPkgsTable {
			// Show Installation key
			if !noEdit {
				guideFrame.AddText("CTRL + D (Install package)", false, tview.AlignRight, tcell.ColorWhite)
				noEdit = true
			}

		} else {
			guideFrame.Clear().
				AddText("F1 (Keys Info) / TAB (navigation)", false, tview.AlignLeft, tcell.ColorWhite)
			noEdit = false
		}

		return event

	})

	// Set focus to the dropdown instead of the flex / Start the event loop
	if err := app.SetRoot(rowFlex, true).SetFocus(packageTable).Run(); err != nil {
		panic(err)
	}

}
