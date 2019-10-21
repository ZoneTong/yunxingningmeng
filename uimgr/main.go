package main

import (
	"fmt"
	_ "image/png"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var (
	loginDialog  *ui.Window
	menuDialog   *ui.Window
	MainTitle    = "运兴合作社管理软件"
	InitPassword = "潼18503021072"
)

func process() {
	ui.OnShouldQuit(func() bool {
		fmt.Println("ui.OnShouldQuit")
		if loginDialog != nil {
			loginDialog.Destroy()
		} else {
			menuDialog.Destroy()
		}
		return true
	})

	// NewLoginDialog()
	// NewMenuDialog()
	// NewChangePasswordDialog()
	// NewTable()
	NewPurchaseTable()
}

func main() {
	ui.Main(process)
}
