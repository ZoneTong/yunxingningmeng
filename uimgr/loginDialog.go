package main

import (
	"fmt"

	"github.com/andlabs/ui"
)

var (
	accountEntry  = ui.NewEntry()
	passwordEntry = ui.NewPasswordEntry()
	passwordBtn   = ui.NewButton("修改密码")
	loginBtn      = ui.NewButton("登陆")
)

func NewLoginDialog() {

	loginDialog = ui.NewWindow(MainTitle, 300, 160, false)
	loginDialog.OnClosing(func(*ui.Window) bool {
		// loginDialog = nil
		fmt.Println("loginDialog OnClosing")
		ui.Quit()
		return true
	})

	grid := ui.NewGrid()
	grid.SetPadded(true)

	grid.Append(ui.NewLabel("账号"), 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(accountEntry, 1, 0, 2, 1, true, ui.AlignFill, false, ui.AlignFill)

	grid.Append(ui.NewLabel("密码"), 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(passwordEntry, 1, 1, 2, 1, true, ui.AlignFill, false, ui.AlignFill)

	grid.Append(passwordBtn, 1, 2, 1, 2, false, ui.AlignFill, true, ui.AlignFill)
	grid.Append(loginBtn, 2, 2, 1, 2, true, ui.AlignFill, true, ui.AlignFill)
	loginBtn.OnClicked(func(*ui.Button) {
		NewMenuDialog()
		loginDialog.Hide()
		loginDialog.Destroy()
		loginDialog = nil
	})

	loginDialog.SetChild(grid)
	loginDialog.SetMargined(true)
	// loginDialog.SetBorderless(true)

	loginDialog.Show()
	accountEntry.SetText("root")
	passwordEntry.SetText("123456")
}
