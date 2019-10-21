package main

import "github.com/andlabs/ui"

type modelHandler struct {
}

func newModelHandler() *modelHandler {
	m := new(modelHandler)

	return m
}

func (mh *modelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""), // column 0 text
		ui.TableString(""), // column 1 text
		ui.TableString(""), // column 2 text
		ui.TableColor{},    // row background color
		ui.TableColor{},    // column 1 text color
		ui.TableImage{},    // column 1 image
		ui.TableString(""), // column 4 button text
		ui.TableInt(0),     // column 3 checkbox state
		ui.TableInt(0),     // column 5 progress
	}
}

func (mh *modelHandler) NumRows(m *ui.TableModel) int {
	return 15
}

func (mh *modelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	return ui.TableString("aat")
}

func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {

}

var (
// tableDialog *ui.Window
)

func NewTable() {
	tableDialog := ui.NewWindow(MainTitle+"-货物采购表", 300, 400, false)
	tableDialog.OnClosing(func(w *ui.Window) bool {
		menuDialog.Show()
		return true
	})

	mh := newModelHandler()
	model := ui.NewTableModel(mh)

	table := ui.NewTable(&ui.TableParams{
		Model:                         model,
		RowBackgroundColorModelColumn: 3,
	})
	tableDialog.SetChild(table)
	tableDialog.SetMargined(true)

	// table.AppendTextColumn("Column 1", 0, ui.TableModelColumnNeverEditable, nil)

	tableDialog.Show()
}
