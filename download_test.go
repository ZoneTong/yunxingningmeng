package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExportExcel(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1", args{[]interface{}{PurchaseRecord{Id: 32}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExportExcel(tt.args.values)
		})
	}
}

func ExportExcel(values []interface{}) {
	if len(values) < 0 {
		return
	}
	// num := ty.NumField()
	// for i := 0; i < num; i++ {
	// 	f := ty.Field(i)
	// 	fmt.Println(f.Name)

	// }

	rv := reflect.Indirect(reflect.ValueOf(values[0]))
	fv := rv.FieldByName("Id")
	fmt.Println(rv.Type().Field(0).Name, fv.Interface())

}
