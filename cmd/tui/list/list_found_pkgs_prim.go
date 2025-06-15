package list

import "github.com/rivo/tview"

func ListFoundPkgs() *tview.Table {

	packageTable := tview.NewTable()

	packageTable.SetBorder(true)
	packageTable.SetTitle("📥Found Packages")

	packageTable.Select(0, 0).
		SetSelectable(true, true)

	return packageTable

}
