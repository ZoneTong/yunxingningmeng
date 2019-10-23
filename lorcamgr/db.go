package main

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

var (
	db orm.Ormer
)

func init() {
	// orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	// orm.ResetModelCache()
	orm.RegisterModelWithPrefix("t_", new(User), new(PurchaseRecord), new(SaleRecord), new(CostRecord), new(StockRecord), new(FinanceStatics))

	db = orm.NewOrm()

}

type User struct {
	Id       int
	Name     string
	Password string

	Deleted int       // 删除标志
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

// type Record struct {
// 	Id   int `orm:"pk;auto;null"`
// 	Date string
// }

// func (r *Record) VerifyDate() bool {
// 	return verifyDate(r.Date) == nil
// }

type PurchaseRecord struct {
	// Record
	Id   int
	Date string

	Provider             string // 供应商
	Logistics            string // 物流
	Specifications       string // 规格
	Number               int    // 件数
	Weight, Price, Total float64

	Deleted int       // 删除标志
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

func (p *PurchaseRecord) CalcTotal() float64 {
	return FloatMul(p.Weight, p.Price)
}

type SaleRecord struct {
	// Record
	Id   int `orm:"pk;auto;null"`
	Date string

	Customer             string // 客户名
	Logistics            string // 物流
	Specifications       string
	Number               int
	Weight, Price, Total float64

	Deleted int       // 删除标志
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

func (p *SaleRecord) CalcTotal() float64 {
	return FloatMul(p.Weight, p.Price)
}

type CostRecord struct {
	// Record
	Id   int `orm:"pk;auto;null"`
	Date string

	TeaCost, LaborCost, Freight, Postage, CartonCost float64
	Consumables, PackingCharges                      float64
	Total                                            float64

	Deleted int       // 删除标志
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

func (r *CostRecord) CalcTotal() float64 {
	return FloatSum(r.TeaCost, r.LaborCost, r.Freight, r.Postage, r.CartonCost, +r.Consumables, +r.PackingCharges)
}

type StockRecord struct {
	// Record
	Id   int `orm:"pk;auto;null"`
	Date string

	Specifications string
	Operation      string
	Weight, Money  float64
	SurplusStock   float64 // 剩余库存

	Deleted int       // 删除标志
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

type FinanceStatics struct {
	// Record
	Id   int `orm:"pk;auto;null"`
	Date string

	Purchase, Sale, Cost, LastBalance, Total float64

	Deleted int       // 删除标志
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
}

func (r *FinanceStatics) CalcTotal() float64 {
	return FloatSum(r.Sale, -r.Purchase, -r.Cost, +r.LastBalance)
}

func FloatSum(floats ...float64) float64 {
	r := decimal.NewFromFloat(0)
	for _, f := range floats {
		r = r.Add(decimal.NewFromFloat(f))
	}
	f, _ := r.Float64()
	return f
}

func FloatMul(floats ...float64) float64 {
	r := decimal.NewFromFloat(1)
	for _, f := range floats {
		r = r.Mul(decimal.NewFromFloat(f))
	}
	f, _ := r.Float64()
	return f
}
