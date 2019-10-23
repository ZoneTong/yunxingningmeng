package main

import (
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func TestDB(t *testing.T) {
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := new(User)
	user.Name = "yunxing"
	err := o.Read(user, "Name")
	if err != nil {
		user.Name = "yunxing"
		user.Password = "750588"
		t.Log(o.Insert(user))
	}

	// r := new(PurchaseRecord)
	// r.Date = "20151121"
	// r.Provider = "gongying"
	// r.Logistics = "wuliu"
	// r.Number = 4
	// r.Weight = 6.5
	// r.Price = 3.2
	// r.Total = r.CalcTotal()
	// t.Log(o.Insert(r))
}

// 精度存在问题
func TestCostRecord_CalcTotal(t *testing.T) {
	type fields struct {
		Id             int
		Date           string
		TeaCost        float64
		LaborCost      float64
		Freight        float64
		Postage        float64
		CartonCost     float64
		Consumables    float64
		PackingCharges float64
		Total          float64
		Created        time.Time
		Updated        time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"1", fields{1, "20191022", 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 8.8, 0, time.Now(), time.Now()}, 31.9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CostRecord{
				Id:             tt.fields.Id,
				Date:           tt.fields.Date,
				TeaCost:        tt.fields.TeaCost,
				LaborCost:      tt.fields.LaborCost,
				Freight:        tt.fields.Freight,
				Postage:        tt.fields.Postage,
				CartonCost:     tt.fields.CartonCost,
				Consumables:    tt.fields.Consumables,
				PackingCharges: tt.fields.PackingCharges,
				Total:          tt.fields.Total,
				Created:        tt.fields.Created,
				Updated:        tt.fields.Updated,
			}
			if got := r.CalcTotal(); got != tt.want {
				t.Errorf("CostRecord.CalcTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
