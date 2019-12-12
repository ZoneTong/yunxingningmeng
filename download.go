package main

import (
	"reflect"

	"github.com/tealeg/xlsx"
)

func saveExcel(path string, rows [][]interface{}, columns []string) (err error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return
	}

	row = sheet.AddRow()
	for _, col := range columns {
		cell = row.AddCell()
		cell.SetValue(col)
	}

	for _, vr := range rows {
		row = sheet.AddRow()
		for _, v := range vr {
			cell = row.AddCell()
			cell.SetValue(v)
		}
	}

	err = file.Save(path)
	if err != nil {
		return
	}

	return
}

func reflectRowsValues(sliceValues interface{}, fieldNames []string) (rows [][]interface{}) {
	vs := reflect.ValueOf(sliceValues)
	if vs.Kind() != reflect.Slice {
		panic("values not slice")
		// return
	}

	if vs.Len() < 0 {
		return
	}

	values := make([]interface{}, vs.Len())
	for i := 0; i < vs.Len(); i++ {
		values[i] = vs.Index(i).Interface()
	}

	rows = make([][]interface{}, 0, len(values))
	for _, v := range values {
		rv := reflect.Indirect(reflect.ValueOf(v))
		row := make([]interface{}, 0, len(fieldNames))
		for _, fn := range fieldNames {
			fv := rv.FieldByName(fn)
			row = append(row, fv.Interface())
		}
		rows = append(rows, row)
	}

	return
}
