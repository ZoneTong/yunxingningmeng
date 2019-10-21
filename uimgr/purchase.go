package main

import (
	"fmt"

	"github.com/andlabs/ui"
)

type purchaseModelHandler struct {
	ui.TableModelHandler
	data  []*PurchaseRecord
	model *ui.TableModel
}

func (h *purchaseModelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{ui.TableString("0"), ui.TableString(""), ui.TableString(""), ui.TableString(""), ui.TableString(""), ui.TableInt(0), ui.TableString("0.0"), ui.TableString("0.0"), ui.TableString("0.0")}
}

func (h *purchaseModelHandler) NumRows(m *ui.TableModel) int {
	return len(h.data)
}

func (h *purchaseModelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	switch column {
	case 0:
		return ui.TableString(fmt.Sprint(h.data[row].Id))
	case 1:
		return ui.TableString(h.data[row].Date)
	case 2:
		return ui.TableString(h.data[row].Provider)
	case 3:
		return ui.TableString(h.data[row].Logistics)
	case 4:
		return ui.TableString(h.data[row].Specifications)
	case 5:
		return ui.TableInt(h.data[row].Number)
	case 6:
		return ui.TableString(fmt.Sprint(h.data[row].Weight))
	case 7:
		return ui.TableString(fmt.Sprint(h.data[row].Price))
	case 8:
		return ui.TableString(fmt.Sprint(h.data[row].Total()))
	}
	return ui.TableString("unkown")
}

func (h *purchaseModelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	switch column {
	case 0:

	}
	m.RowChanged(row)
}

func (h *purchaseModelHandler) RefreshData() {
	for i := 0; i < len(h.data); i++ {
		h.model.RowChanged(i)
	}
	// m.RowChanged(row)
}

func NewPurchaseTable() {
	purchaseHandle := new(purchaseModelHandler)
	purchaseHandle.data = []*PurchaseRecord{&PurchaseRecord{Id: 1, Date: "20191014"}}

	purchaseModel := ui.NewTableModel(purchaseHandle)
	purchaseHandle.model = purchaseModel
	purchaseTable := ui.NewTable(
		&ui.TableParams{Model: purchaseModel},
	)
	purchaseTable.AppendTextColumn("序号", 0, ui.TableModelColumnNeverEditable, nil)
	purchaseTable.AppendTextColumn("日期", 1, ui.TableModelColumnNeverEditable, nil)
	purchaseTable.AppendTextColumn("净重(斤)", 6, ui.TableModelColumnNeverEditable, nil)

	purchaseTable.AppendImageColumn("删除", 8)
	tableDialog := ui.NewWindow(MainTitle+"-货物采购表", 300, 400, false)
	tableDialog.OnClosing(func(w *ui.Window) bool {
		menuDialog.Show()
		return true
	})

	tableDialog.SetChild(purchaseTable)
	tableDialog.SetMargined(true)

	purchaseHandle.data[0].Weight = 0.3
	tableDialog.Show()

	// 改变数据
	purchaseHandle.data[0].Weight = 0.6
	purchaseHandle.RefreshData()

	// 没找到行点击事件
	purchaseHandle.data = append(purchaseHandle.data, &PurchaseRecord{Id: 2, Date: "20190304", Weight: 0.9})
	purchaseHandle.model.RowInserted(1)

}

func NewPurchaseDialog() {

}
