package list

import (
	"depoty/internal/listing"
	"runtime"
	"strings"

	"github.com/rivo/tview"
)

func RefreshWholePkgs(packageTable *tview.Table, listOfPkgs *map[string]string) []string {
	for k := range *listOfPkgs {
		delete(*listOfPkgs, k)
	}

	outLit := make([]string, 0)

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := listing.ListPkgs()
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	switch runtime.GOOS {
	case "windows":
		if len(pkgs)-1 == 0 {
			return nil
		} else {
			// Row Counter
			j := 0

			// Packges Slice
			// Create the packageTable
			for i := 0; i < len(pkgs)-1; i++ {
				outLit = append(outLit, pkgs[i])
				(*listOfPkgs)[strings.Split(pkgs[i], " ")[0]] = pkgs[i]
				j++
			}

			return outLit
		}
	case "linux":

		j := 0

		for i := 0; i < len(pkgs); i++ {
			outLit = append(outLit, pkgs[i])
			(*listOfPkgs)[strings.Split(pkgs[i], " ")[0]] = pkgs[i]
			j++
		}
		return outLit
	default:
		return make([]string, 0)
	}

}

func ListWholePkgs(packageTable *tview.Table) map[string]string {

	list := make(map[string]string)

	pkgChan := make(chan []string, 1)

	go func() {
		pkgs := listing.ListPkgs()
		pkgChan <- pkgs
		close(pkgChan)
	}()

	pkgs := <-pkgChan

	packageTable.SetBorder(true)
	packageTable.SetTitle("📦Installed Packages")

	packageTable.Select(0, 0).
		SetSelectable(true, true)

	switch runtime.GOOS {
	case "windows":
		if len(pkgs)-1 == 0 {
			return nil
		} else {
			// Row Counter
			j := 0

			// Packges Slice
			// Create the packageTable
			// -1 because we have an empty line (\n) at the end
			for i := 0; i < len(pkgs)-1; i++ {
				packageTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
				list[strings.Split(pkgs[i], " ")[0]] = pkgs[i]
				j++
			}

			return list
		}
	case "linux":
		// Row Counter
		j := 0

		// Packges Slice
		// Create the packageTable
		for i := 0; i < len(pkgs); i++ {
			packageTable.SetCell(j, 0, tview.NewTableCell(pkgs[i]))
			list[strings.Split(pkgs[i], " ")[0]] = pkgs[i]
			j++
		}

		return list
	default:
		return make(map[string]string)
	}

}
