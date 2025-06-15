package tui

import "github.com/rivo/tview"

func guideInfoModalPage() *tview.Pages {
	KeyDataModal := tview.NewModal().SetText(`
		
		Navigation Keys:
		CTRL + S (for search)
		CTRL + F (Found Packages)
		CTRL + P (Installed Packges)
		CTRL + N (Package Info)

		General keys:
		CTRL + R (Refresh packages & viewport)
		Esc (Back)
		
	`)

	KeyDataModal.SetTitle("Keys Information")

	// Show Keys information
	return tview.NewPages().AddPage("Keys Information", KeyDataModal, false, true)

}
