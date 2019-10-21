package main

import (
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func TestDB(t *testing.T) {
	orm.RunSyncdb("default", true, true)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := new(User)
	user.Name = "yunxing"
	user.Password = "750588"
	// err := o.Read(user, "Name")
	// if

	t.Log(o.Insert(user))

	r := new(PurchaseRecord)
	r.Date = "20151121"
	r.Provider = "gongying"
	r.Logistics = "wuliu"
	r.Number = 4
	r.Weight = 6.5
	r.Price = 3.2
	r.Total = r.CalcTotal()
	t.Log(o.Insert(r))
}
