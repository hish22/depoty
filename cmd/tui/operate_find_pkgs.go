package tui

import (
	"depoty/internal/finding"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func operateFindingPkgs(installPkg *tview.InputField, PkgsTable *tview.Table) {
	// Clear the found packages content
	PkgsTable.Clear()

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
				PkgsTable.SetBorderColor(tcell.ColorRed)
				PkgsTable.SetTitle(pkgs[2])
				break
			}
		}

		// Get rid of unwanted choco ads (Check this)
		if strings.Contains(pkgs[len(pkgs)-3], "Learn more") && i == len(pkgs)-3 {
			PkgsTable.SetTitle(pkgs[i-3])
			PkgsTable.SetBorderColor(tcell.ColorGreen)
			break
		} else {
			if len(pkgs)-2 == i {
				PkgsTable.SetTitle(pkgs[i])
				PkgsTable.SetBorderColor(tcell.ColorGreen)
				break
			}
		}

		PkgsTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
		j++
	}
}
