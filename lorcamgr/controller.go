package main

import (
	"database/sql"
	"log"
)

func verifyPassword(account, password string) string {
	u := new(User)
	u.Name = account
	err := db.Read(u, "Name")
	if err != nil {
		if err == sql.ErrNoRows {
			return "账号不存在"
		}
		return err.Error()
	}

	// "yunxing"
	//  "750588"
	if u.Password == password {
		return "ok"
	}

	return "账号或密码错误"
}

func NewOrEditPurchaseRecord(r *PurchaseRecord) string {
	log.Println("NewOrEditPurchaseRecord", *r)
	var err error
	r.Total = r.CalcTotal()
	if r.Id == 0 {
		_, err = db.Insert(r)
	} else {
		_, err = db.Update(r)
	}

	if err != nil {
		return err.Error()
	}
	return "ok"
}

func NewOrEditSaleRecord(r *SaleRecord) string {
	log.Println("NewOrEditSaleRecord", *r)
	var err error
	r.Total = r.CalcTotal()
	if r.Id == 0 {
		_, err = db.Insert(r)
	} else {
		_, err = db.Update(r)
	}

	if err != nil {
		return err.Error()
	}
	return "ok"
}

const (
	SEARCH_DATE = iota
	SEARCH_PROVIDER
	SEARCH_CUSTOMER
)

type Response struct {
	Data  interface{}
	Error string
}

func SearchPurchaseRecords(column int, key string) (resp *Response) {
	resp = new(Response)
	table := db.QueryTable(&PurchaseRecord{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	} else {
		table = table.Filter("provider__icontains", key)
	}

	var rows []*PurchaseRecord
	n, err := table.All(&rows)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	log.Println("SearchPurchaseRecords Total count:", n)
	resp.Data = rows
	return
}

func SearchSaleRecords(column int, key string) (resp *Response) {
	resp = new(Response)
	table := db.QueryTable(&SaleRecord{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	} else {
		table = table.Filter("customer__icontains", key)
	}

	var rows []*SaleRecord
	n, err := table.All(&rows)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	log.Println("SearchSaleRecords Total count:", n)
	resp.Data = rows
	return
}

func DelPurchaseRecord(id int) string {
	r := PurchaseRecord{Id: id}
	_, err := db.Delete(&r)
	defer log.Println("DelPurchaseRecord", id, err)
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func DelSaleRecord(id int) string {
	r := SaleRecord{Id: id}
	_, err := db.Delete(&r)
	defer log.Println("DelSaleRecord", id, err)
	if err != nil {
		return err.Error()
	}
	return "ok"
}
