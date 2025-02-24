package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func searchPkgs() (*tview.Flex, *tview.InputField, *tview.TextView) {
	//Search bar components to install package
	searchPkg := tview.NewInputField()

	searchPkg.SetLabel("Find Package:")

	searchPkg.SetFieldBackgroundColor(tcell.ColorYellow)

	searchPkg.SetFieldTextColor(tcell.ColorBlack)

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
