package tui

import (
	"depoty/internal/finding"
	"depoty/internal/listing"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ListWholePkgs() *tview.Table {
	pkgsTable := tview.NewTable()

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := listing.ListPkgs()
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	// Row Counter
	j := 0

	// Create the pkgsTable
	for i := 0; i < len(pkgs); i += 3 {
		pkgsTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
		j++
	}

	pkgsTable.SetBorder(true)
	pkgsTable.SetTitle("Packages")

	pkgsTable.Select(0, 0).
		SetSelectable(true, true)

	return pkgsTable
}

func searchPkgs() *tview.InputField {
	//Search bar components to install package
	installPkg := tview.NewInputField()

	installPkg.SetLabel("Find Package:")

	installPkg.SetBorder(true).SetTitle("Search")

	return installPkg

}

func ListFoundPkgs() *tview.Table {

	pkgsTable := tview.NewTable()

	pkgsTable.SetBorder(true)
	pkgsTable.SetTitle("Found Packages")

	pkgsTable.Select(0, 0).
		SetSelectable(true, true)

	return pkgsTable

}

func TuiStart() {
	// box := tview.NewBox().SetBorder(true).SetTitle("Depoty")
	app := tview.NewApplication()

	pkgsTable := ListWholePkgs()

	installPkg := searchPkgs()

	foundPkgsTable := ListFoundPkgs()

	// Search and packages Flex
	pkgsFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(installPkg, 0, 1, false).
		AddItem(foundPkgsTable, 0, 2, false).
		AddItem(pkgsTable, 0, 2, false)

	// Display the information of a package
	pkgInfo := tview.NewTextView().SetText("Press Enter to fetch a package information.")

	pkgInfo.SetBorder(true).SetTitle("Package Information")

	pkgsTable.SetSelectedFunc(func(row, column int) {
		// Get selected Row/Columns
		r, c := pkgsTable.GetSelection()

		// Seperate name from version
		textSlice := strings.Split(pkgsTable.GetCell(r, c).Text, " ")

		// Print starting fetching
		pkgInfo.SetText("Fetching information...")

		// Start Fetching app/service information
		info := finding.FindPkgInfo(textSlice[0])

		pkgInfo.SetText(info)
	})

	installPkg.SetDoneFunc(func(key tcell.Key) {

		pkgChan := make(chan []string, 1)

		go func() {
			pkgs := finding.FindPkg(installPkg.GetText())
			pkgChan <- pkgs
			close(pkgChan)
		}()

		pkgs := <-pkgChan

		j := 0

		for i := 2; i < len(pkgs); i++ {

			if len(pkgs)-2 == i {
				foundPkgsTable.SetTitle(pkgs[i])
				break
			}
			foundPkgsTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
			j++
		}
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlS {
			app.SetFocus(installPkg)
		} else if event.Key() == tcell.KeyCtrlP {
			app.SetFocus(pkgsTable)
		} else if event.Key() == tcell.KeyTAB {
			app.SetFocus(pkgInfo)
		} else if event.Key() == tcell.KeyCtrlF {
			app.SetFocus(foundPkgsTable)
		}
		return event
	})

	// Create the flex container
	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn). // Explicitly set direction (optional, defaults to Column)
		AddItem(pkgsFlex, 0, 1, true).
		AddItem(pkgInfo, 0, 2, true)

	flex.SetBackgroundColor(tcell.Color102)
	// Set focus to the dropdown instead of the flex
	if err := app.SetRoot(flex, true).SetFocus(pkgsTable).Run(); err != nil {
		panic(err)
	}

}
