package main

import "github.com/andlabs/ui"

func NewMenuDialog() {
	menuDialog = ui.NewWindow(MainTitle+"-功能列表", 640, 480, true)
	menuDialog.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		menuDialog.Destroy()
		return true
	})

	purchaseBtn := ui.NewButton("货物采购")
	purchaseBtn.OnClicked(func(*ui.Button) {
		NewTable()
		menuDialog.Hide()
	})

	saleBtn := ui.NewButton("货物销售")
	saleBtn.OnClicked(func(*ui.Button) {
		NewTable()
		menuDialog.Hide()
	})

	costBtn := ui.NewButton("运营成本")
	costBtn.OnClicked(func(*ui.Button) {
		NewTable()
		menuDialog.Hide()
	})

	stockBtn := ui.NewButton("货物库存")
	stockBtn.OnClicked(func(*ui.Button) {
		NewTable()
		menuDialog.Hide()
	})

	financeBtn := ui.NewButton("财务统计")
	financeBtn.OnClicked(func(*ui.Button) {
		NewTable()
		menuDialog.Hide()
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	vbox.Append(purchaseBtn, false)
	vbox.Append(saleBtn, false)
	vbox.Append(costBtn, false)
	vbox.Append(stockBtn, false)
	vbox.Append(financeBtn, false)
	menuDialog.SetChild(vbox)
	menuDialog.SetMargined(true)

	menuDialog.Show()
}

func NewSaleDialog() {

}

func NewCostDialog() {

}

func NewStockDialog() {

}

func NewFinanceSearchDialog() {

}
