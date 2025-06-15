package operation

import (
	"depoty/internal/finding"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func OperateFindingPkgs(installPkg *tview.InputField, PkgsTable *tview.Table) {
	// Clear the found packages content
	PkgsTable.Clear()

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := finding.FindPkg(installPkg.GetText())
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	// Check if no packages found, else start iterating through these packages
	if len(pkgs)-1 == 0 {
		PkgsTable.SetBorderColor(tcell.ColorRed)
		PkgsTable.SetTitle("0 packages found.")
	} else {
		j := 0
		// Total number of packages in string type
		numberOfPkgs := strconv.Itoa(len(pkgs) - 1)

		// Set the border green color since we found packges to list
		PkgsTable.SetBorderColor(tcell.ColorGreen)
		PkgsTable.SetTitle(numberOfPkgs + " packages found.")
		for i := 0; i < len(pkgs); i++ {

			// Replace the | with space " "
			finalPkg := strings.Replace(pkgs[i], "|", " ", 1)

			// break if we reached the last element
			if len(pkgs)-1 == i {
				break
			}
			// Populate the table with found packages
			PkgsTable.SetCell(j, 0, tview.NewTableCell(finalPkg))
			j++
		}
	}

}
