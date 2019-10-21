package main

import (
	"github.com/andlabs/ui"
)

func NewChangePasswordDialog() {
	dialog := ui.NewWindow(MainTitle, 200, 160, false)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	label_hspan := 2
	grid.Append(ui.NewLabel("旧密码"), 0, 0, label_hspan, 1, false, ui.AlignFill, false, ui.AlignFill)
	oldpasswordBtn := ui.NewPasswordEntry()
	grid.Append(oldpasswordBtn, label_hspan, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(ui.NewLabel("新密码"), 0, 1, label_hspan, 1, false, ui.AlignFill, false, ui.AlignFill)
	newpasswordEntry := ui.NewPasswordEntry()
	grid.Append(newpasswordEntry, label_hspan, 1, label_hspan, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(ui.NewLabel("确认新密码"), 0, 2, label_hspan, 1, false, ui.AlignFill, false, ui.AlignFill)
	confirmPasswordEntry := ui.NewPasswordEntry()
	grid.Append(confirmPasswordEntry, label_hspan, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	okBtn := ui.NewButton("确定")
	grid.Append(okBtn, label_hspan, 3, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	dialog.SetChild(grid)
	dialog.SetMargined(true)
	dialog.Show()
}
