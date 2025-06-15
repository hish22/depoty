package fetch

import (
	"depoty/internal/finding"
	"strings"

	"github.com/rivo/tview"
)

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
