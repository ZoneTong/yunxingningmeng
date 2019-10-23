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

func SearchPurchaseRecords(c SearchCondition) (resp *Response) {
	table := db.QueryTable(&PurchaseRecord{})

	if c.SearchArea == SEARCH_DATE {
		table = table.Filter("date__icontains", c.SearchKey)
	} else {
		table = table.Filter("provider__icontains", c.SearchKey)
	}

	var rows []*PurchaseRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		resp.Error = err.Error()
	}

	_, err = table.All(&rows)
	if err != nil {
		return
	}

	resp.Rows = rows
	log.Println("SearchPurchaseRecords Total count:", resp.Total, len(rows))
	if resp.Total > 0 {
		log.Println("SearchPurchaseRecords first:", rows[0])
	}
	return
}

func searchPartial(table orm.QuerySeter, c SearchCondition) (newtable orm.QuerySeter, resp *Response, err error) {
	log.Println("searchPartial", c)
	resp = new(Response)
	orders := []string{"-date", "-id"}
	newtable = table.Filter("deleted", 0)
	resp.Total, err = newtable.Count()
	// resp.TotalNotFiltered = resp.Total
	if err != nil {
		return
	}

	// orderby
	if c.SortName != "" {
		orders[0] = structFieldName2DBCloumnName(c.SortName)
		if c.SortOrder == "desc" {
			orders[0] = "-" + orders[0]
		}
	}

	log.Println("searchPartial orders", orders)
	newtable = newtable.OrderBy(orders...)

	if c.PageSize > 0 {
		newtable = newtable.Limit(c.PageSize, (c.PageNumber-1)*c.PageSize)
	}
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

func SearchSaleRecords(c SearchCondition) (resp *Response) {
	table := db.QueryTable(&SaleRecord{})
	if c.SearchArea == SEARCH_DATE {
		table = table.Filter("date__icontains", c.SearchKey)
	} else {
		table = table.Filter("customer__icontains", c.SearchKey)
	}

	var rows []*SaleRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		resp.Error = err.Error()
	}

	_, err = table.All(&rows)
	if err != nil {
		return
	}

	resp.Rows = rows

	log.Println("SearchSaleRecords Total count:", resp.Total, len(rows))
	if resp.Total > 0 {
		log.Println("SearchSaleRecords first:", rows[0])
	}
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

func SearchCostRecords(c SearchCondition) (resp *Response) {
	table := db.QueryTable(&CostRecord{})
	if c.SearchArea == SEARCH_DATE {
		table = table.Filter("date__icontains", c.SearchKey)
	}

	var rows []*CostRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		resp.Error = err.Error()
	}

	_, err = table.All(&rows)
	if err != nil {
		return
	}

	resp.Rows = rows

	log.Println("SearchCostRecords Total count:", resp.Total, len(rows))
	if resp.Total > 0 {
		log.Println("SearchCostRecords first:", rows[0])
	}
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

func SearchStockRecords(c SearchCondition) (resp *Response) {
	table := db.QueryTable(&StockRecord{})
	if c.SearchArea == SEARCH_DATE {
		table = table.Filter("date__icontains", c.SearchKey)
	}

	var rows []*StockRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		resp.Error = err.Error()
	}

	_, err = table.All(&rows)
	if err != nil {
		return
	}

	resp.Rows = rows

	log.Println("SearchStockRecords Total count:", resp.Total, len(rows))
	if resp.Total > 0 {
		log.Println("SearchStockRecords first:", rows[0])
	}
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
	table := db.QueryTable(&FinanceStatics{})
	if c.SearchArea == SEARCH_DATE {
		table = table.Filter("date__icontains", c.SearchKey)
	}

	var rows []*FinanceStatics
	table, resp, err := searchPartial(table, c)
	if err != nil {
		resp.Error = err.Error()
	}

	_, err = table.All(&rows)
	if err != nil {
		return
	}

	resp.Rows = rows

	log.Println("SearchFinanceStaticss Total count:", resp.Total, len(rows))
	if resp.Total > 0 {
		log.Println("SearchFinanceStaticss first:", rows[0])
	}
	return
}

func CalcFinanceByDate(date string) string {
	r := FinanceStatics{Date: date}
	defer log.Println("CalcFinanceByDate", r)
	var resp *Response
	c := SearchCondition{SearchArea: SEARCH_DATE, SearchKey: date}
	resp = SearchPurchaseRecords(c)
	for _, row := range resp.Rows.([]*PurchaseRecord) {
		r.Purchase += row.Total
		r.PurchasedStock += row.Weight
	}

	resp = SearchSaleRecords(c)
	for _, row := range resp.Rows.([]*SaleRecord) {
		r.Sale += row.Total
		r.SaledStock += row.Weight
	}

	resp = SearchCostRecords(c)
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
