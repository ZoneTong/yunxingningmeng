package main

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(new(User), new(PurchaseRecord), new(SaleRecord), new(CostRecord), new(StockRecord))
}

type User struct {
	Id   int
	Name string
	// Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	// Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
	Password   string
	ModifyTime time.Time
}

// type Record struct {
// 	Id   int
// 	Date string
// }

// func (r *Record) VerifyDate() bool {
// 	return verifyDate(r.Date) == nil
// }

type PurchaseRecord struct {
	// Record
	Id   int
	Date string

	Provider       string // 供应商
	Logistics      string // 物流
	Specifications string // 规格
	Number         int    // 件数
	Weight, Price  float64
}

func (p *PurchaseRecord) Total() float64 {
	return p.Weight * p.Price
}

func (p *PurchaseRecord) SetPrice(c float64) {
	p.Price = c
}

type SaleRecord struct {
	// Record
	Id   int
	Date string

	Customer       string // 供应商
	Logistics      string // 物流
	Specifications string
	Number         int
	Weight, Price  float64
}

func (p *SaleRecord) Total() float64 {
	return p.Weight * p.Price
}

type CostRecord struct {
	// Record
	Id   int
	Date string
	string
	TeaCost, LaborCost, Freight, Postage, CartonCost, Consumables, PackingCharges float64
}

func (r *CostRecord) Total() float64 {
	return r.TeaCost + r.LaborCost + r.Freight + r.Postage + r.CartonCost + r.Consumables + r.PackingCharges
}

type StockRecord struct {
	// Record
	Id   int
	Date string

	Specifications string
	Operation      string
	Weight, Money  float64
	SurplusStock   float64 // 剩余库存
}

type FinanceStatics struct {
	// Record
	Id   int
	Date string

	Purchase, Sale, Cost, LastBalance float64
}

func (r *FinanceStatics) Total() float64 {
	return r.Sale - r.Purchase - r.Cost + r.LastBalance
}
