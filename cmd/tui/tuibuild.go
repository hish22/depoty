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
	pkgsTable.SetTitle("Installed Packages")

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

func startFindOper(pkgInfo *tview.TextView, pkgsTable *tview.Table) {
	/* These two step are not working, check them then */
	// Clearing text buffer
	pkgInfo.Clear()

	// Print starting fetching
	pkgInfo.SetText("Fetching information...").SetDynamicColors(true)
	// Get selected Row/Columns
	r, c := pkgsTable.GetSelection()

	// Seperate name from version
	textSlice := strings.Split(pkgsTable.GetCell(r, c).Text, " ")

	// Start Fetching app/service information
	info := finding.FindPkgInfo(textSlice[0])

	infoSplit := strings.Split(info, "\n")

	var resultInfo strings.Builder

	for i, info := range infoSplit {
		if i < 18 && i >= 3 {
			resultInfo.WriteString("[yellow]" + info + "[-]\n")
		} else {
			resultInfo.WriteString(info + "\n")
		}
	}

	pkgInfo.SetText(resultInfo.String())

}

func operateFoundPkgs(installPkg *tview.InputField, foundPkgsTable *tview.Table) {
	foundPkgsTable.Clear()

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := finding.FindPkg(installPkg.GetText())
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	j := 0

	for i := 2; i < len(pkgs); i++ {

		if len(pkgs) == 4 || len(pkgs) == 8 {
			if strings.Contains(pkgs[2], "0 packages found.") {
				foundPkgsTable.SetBorderColor(tcell.ColorRed)
				foundPkgsTable.SetTitle(pkgs[2])
				break
			}
		}

		if len(pkgs)-2 == i {
			foundPkgsTable.SetTitle(pkgs[i])
			foundPkgsTable.SetBorderColor(tcell.ColorGreen)
			break
		}
		foundPkgsTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
		j++
	}
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

	// fetch the installed package information
	pkgsTable.SetSelectedFunc(func(row, column int) {
		startFindOper(pkgInfo, pkgsTable)
	})

	// Fetch the found package information.
	foundPkgsTable.SetSelectedFunc(func(row, column int) {
		startFindOper(pkgInfo, foundPkgsTable)
	})

	installPkg.SetDoneFunc(func(key tcell.Key) {
		operateFoundPkgs(installPkg, foundPkgsTable)
	})

	// Guide Info components
	guideFrame := tview.NewFrame(tview.NewBox()).
		AddText("CTRL + S (for search) / CTRL + F (Found Packages) / CTRL + P (Installed Packges)", false, tview.AlignLeft, tcell.ColorWhite).
		AddText("Enter (action)", false, tview.AlignCenter, tcell.ColorWhite)

	// Change the components focus
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
		// Change the frame text based on focusing components
		if app.GetFocus() == foundPkgsTable {
			guideFrame.AddText("CTRL + i (Install package)", false, tview.AlignRight, tcell.ColorWhite)
		} else {
			guideFrame.Clear().
				AddText("CTRL + S (for search) / CTRL + F (Found Packages) / CTRL + P (Installed Packges)", false, tview.AlignLeft, tcell.ColorWhite).
				AddText("Enter (action)", false, tview.AlignCenter, tcell.ColorWhite)
		}

		return event

	})

	// Create the flex container
	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn). // Explicitly set direction (optional, defaults to Column)
		AddItem(pkgsFlex, 0, 1, false).
		AddItem(pkgInfo, 0, 2, false)

	rowFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 12, false).
		AddItem(guideFrame, 0, 1, false)

	flex.SetBackgroundColor(tcell.Color102)
	// Set focus to the dropdown instead of the flex
	if err := app.SetRoot(rowFlex, true).SetFocus(pkgsTable).Run(); err != nil {
		panic(err)
	}

}
