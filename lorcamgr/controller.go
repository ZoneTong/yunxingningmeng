package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
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

func SearchPurchaseRecords(column int, key string) (resp *Response) {
	resp = new(Response)
	table := db.QueryTable(&PurchaseRecord{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	} else {
		table = table.Filter("provider__icontains", key)
	}

	var rows []*PurchaseRecord
	n, err := table.Filter("deleted", 0).OrderBy("-date", "-id").All(&rows)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	log.Println("SearchPurchaseRecords Total count:", n)
	if n > 0 {
		log.Println("SearchPurchaseRecords first:", rows[0])
	}
	resp.Rows = rows
	return
}

func DelPurchaseRecord(id int) string {
	r := PurchaseRecord{Id: id}
	// _, err := db.Delete(&r)
	r.Deleted = 1
	_, err := db.Update(&r, "Deleted")
	defer log.Println("DelPurchaseRecord", id, err)
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

func SearchSaleRecords(column int, key string) (resp *Response) {
	resp = new(Response)
	table := db.QueryTable(&SaleRecord{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	} else {
		table = table.Filter("customer__icontains", key)
	}

	var rows []*SaleRecord
	n, err := table.Filter("deleted", 0).OrderBy("-date", "-id").All(&rows)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	log.Println("SearchSaleRecords Total count:", n)
	if n > 0 {
		log.Println("SearchSaleRecords first:", rows[0])
	}
	resp.Rows = rows
	return
}

func DelSaleRecord(id int) string {
	r := SaleRecord{Id: id}
	// _, err := db.Delete(&r)
	r.Deleted = 1
	_, err := db.Update(&r, "Deleted")
	defer log.Println("DelSaleRecord", id, err)
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func NewOrEditCostRecord(r *CostRecord) string {
	log.Println("NewOrEditCostRecord", *r)
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

func SearchCostRecords(column int, key string) (resp *Response) {
	resp = new(Response)
	table := db.QueryTable(&CostRecord{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	}

	var rows []*CostRecord
	n, err := table.Filter("deleted", 0).OrderBy("-date", "-id").All(&rows)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	log.Println("SearchCostRecords Total count:", n)
	if n > 0 {
		log.Println("SearchCostRecords first:", rows[0])
	}
	resp.Rows = rows
	return
}

func DelCostRecord(id int) string {
	r := CostRecord{Id: id}
	// _, err := db.Delete(&r)
	r.Deleted = 1
	_, err := db.Update(&r, "Deleted")
	defer log.Println("DelCostRecord", id, err)
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func NewOrEditStockRecord(r *StockRecord) string {
	log.Println("NewOrEditStockRecord", *r)
	var err error
	// r.Total = r.CalcTotal()
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

func SearchStockRecords(column int, key string) (resp *Response) {
	resp = new(Response)
	table := db.QueryTable(&StockRecord{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	}

	var rows []*StockRecord
	n, err := table.Filter("deleted", 0).OrderBy("-date", "-id").All(&rows)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	log.Println("SearchStockRecords Total count:", n)
	if n > 0 {
		log.Println("SearchStockRecords first:", rows[0])
	}
	resp.Rows = rows
	return
}

func DelStockRecord(id int) string {
	r := StockRecord{Id: id}
	// _, err := db.Delete(&r)
	r.Deleted = 1
	_, err := db.Update(&r, "Deleted")
	defer log.Println("DelStockRecord", id, err)
	if err != nil {
		return err.Error()
	}
	return "ok"
}

func newOrEditFinanceStatics(r *FinanceStatics) string {
	log.Println("NewOrEditFinanceStatics", *r)
	var err error
	r.Total = r.CalcTotal()
	r.TotalStock = r.CalcStock()
	old := FinanceStatics{Date: r.Date}
	_, _, err = db.ReadOrCreate(&old, "Date")
	if err != nil {
		return err.Error()
	}

	r.Id = old.Id
	_, err = db.Update(r)
	if err != nil {
		return err.Error()
	}

	return "ok"
}

func SearchFinanceStaticss(c SearchCondition) (resp *Response) {
	log.Println("SearchFinanceStaticss", c)
	column := c.SearchArea
	key := c.SearchKey

	resp = new(Response)
	table := db.QueryTable(&FinanceStatics{})

	if column == SEARCH_DATE {
		table = table.Filter("date__icontains", key)
	}

	var rows []*FinanceStatics
	var err error
	orders := []string{"-date", "-id"}
	table = table.Filter("deleted", 0)
	resp.Total, err = table.Count()
	// resp.TotalNotFiltered = resp.Total
	if err != nil {
		goto ERR
	}

	// orderby
	if c.SortName != "" {
		orders[0] = structFieldName2DBCloumnName(c.SortName)
		if c.SortOrder == "desc" {
			orders[0] = "-" + orders[0]
		}
	}
	table = table.OrderBy(orders...)

	_, err = table.Limit(c.PageSize, (c.PageNumber-1)*c.PageSize).All(&rows)
	if err != nil {
		goto ERR
	}
	log.Println("SearchFinanceStaticss Total count:", resp.Total, len(rows))
	if resp.Total > 0 {
		log.Println("SearchFinanceStaticss first:", rows[0])
	}
	resp.Rows = rows
	return

ERR:
	if err != nil {
		resp.Error = err.Error()
	}
	return
}

func CalcFinanceByDate(date string) string {
	r := FinanceStatics{Date: date}
	defer log.Println("CalcFinanceByDate", r)
	var resp *Response
	resp = SearchPurchaseRecords(SEARCH_DATE, date)
	for _, row := range resp.Rows.([]*PurchaseRecord) {
		r.Purchase += row.Total
		r.PurchasedStock += row.Weight
	}

	resp = SearchSaleRecords(SEARCH_DATE, date)
	for _, row := range resp.Rows.([]*SaleRecord) {
		r.Sale += row.Total
		r.SaledStock += row.Weight
	}

	resp = SearchCostRecords(SEARCH_DATE, date)
	for _, row := range resp.Rows.([]*CostRecord) {
		r.Cost += row.Total
	}

	predate, err := nextNDate(date, -1)
	if err != nil {
		return err.Error()
	}

	var pre_f = FinanceStatics{Date: predate}
	err = db.Read(&pre_f, "Date")
	// log.Println("FinanceStatics Read", err)
	if err != nil && err != orm.ErrNoRows {
		return err.Error()
	}
	r.LastBalance = pre_f.Total
	r.LastStock = pre_f.TotalStock

	return newOrEditFinanceStatics(&r)
}

const (
	DATE_FORMAT = "20060102"
)

func CalcFinanceByPeriod(from, to string) (errstr string) {
	errstr = "ok"
	today := time.Now().Local().Format(DATE_FORMAT)
	var date = from
	var err error
	for date <= today && date <= to {
		errstr = CalcFinanceByDate(date)
		if errstr != "ok" {
			break
		}

		date, err = nextNDate(date, 1)
		if err != nil {
			errstr = err.Error()
			break
		}
	}

	return
}
