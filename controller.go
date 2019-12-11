package main

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	SEARCH_MAX_COUNT = 10000
	INIT_MIN_DATE    = "99991229"
)

func verifyPassword(account, password string) string {
	if hoursleft < 0 {
		return fmt.Sprintf("试用期已过,请联系%s获取正式版", author)
	}

	u := new(User)
	u.Name = account
	err := dborm.Read(u, "Name")
	if err != nil {
		if err == orm.ErrNoRows {
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
		_, err = dborm.Insert(r)
	} else {
		_, err = dborm.Update(r)
	}

	if err != nil {
		return err.Error()
	}
	return "ok"
}

func SearchPurchaseRecords(c SearchCondition) (resp *Response) {
	table := dborm.QueryTable(&PurchaseRecord{})

	if c.SearchKeys[0] != "" {
		table = table.Filter("date__icontains", c.SearchKeys[0])
	}
	if c.SearchKeys[1] != "" {
		table = table.Filter("provider__icontains", c.SearchKeys[1])
	}

	var static PurchaseRecord
	min_date := INIT_MIN_DATE
	var max_date string

	var rows []*PurchaseRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		goto ERR
	}

	if resp.Total == 0 {
		return
	}

	_, err = table.All(&rows)
	if err != nil {
		goto ERR
	}

	resp.Rows = rows
	log.Println("SearchPurchaseRecords Total count:", resp.Total, len(rows), rows[0])

	// 统计
	if c.nostatics {
		return
	}
	if resp.Total < SEARCH_MAX_COUNT && resp.Total > len(rows) {
		_, err = table.Limit(SEARCH_MAX_COUNT, 0).All(&rows)
		if err != nil {
			goto ERR
		}
	}
	for _, row := range rows {
		if min_date > row.Date {
			min_date = row.Date
		}
		if max_date < row.Date {
			max_date = row.Date
		}
		static.Number += row.Number
		static.Weight += row.Weight
		static.Total += row.Total
	}
	static.Date = min_date + "-<br/>" + max_date
	static.Provider = c.SearchKeys[1]
	resp.Static = &static
	return
ERR:
	resp.Error = err.Error()
	return
}

func searchPartial(table orm.QuerySeter, c SearchCondition) (newtable orm.QuerySeter, resp *Response, err error) {
	log.Println("searchPartial", c)
	resp = new(Response)
	orders := []string{"-date", "-id"}
	newtable = table.Filter("deleted", 0)
	total, err := newtable.Count()
	resp.Total = int(total)
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
	// _, err := dborm.Delete(&r)
	r.Deleted = 1
	_, err := dborm.Update(&r, "Deleted")
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
		_, err = dborm.Insert(r)
	} else {
		_, err = dborm.Update(r)
	}

	if err != nil {
		return err.Error()
	}
	return "ok"
}

func SearchSaleRecords(c SearchCondition) (resp *Response) {
	table := dborm.QueryTable(&SaleRecord{})
	if c.SearchKeys[0] != "" {
		table = table.Filter("date__icontains", c.SearchKeys[0])
	}
	if c.SearchKeys[1] != "" {
		table = table.Filter("customer__icontains", c.SearchKeys[1])
	}

	var static SaleRecord
	min_date := INIT_MIN_DATE
	var max_date string

	var rows []*SaleRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		goto ERR
	}

	if resp.Total == 0 {
		return
	}

	_, err = table.All(&rows)
	if err != nil {
		goto ERR
	}

	resp.Rows = rows
	log.Println("SearchSaleRecords Total count:", resp.Total, len(rows), rows[0])

	// 统计
	if c.nostatics {
		return
	}
	if resp.Total < SEARCH_MAX_COUNT && resp.Total > len(rows) {
		_, err = table.Limit(SEARCH_MAX_COUNT, 0).All(&rows)
		if err != nil {
			goto ERR
		}
	}
	for _, row := range rows {
		if min_date > row.Date {
			min_date = row.Date
		}
		if max_date < row.Date {
			max_date = row.Date
		}
		static.Number += row.Number
		static.Weight += row.Weight
		static.Total += row.Total
	}
	static.Date = min_date + "-<br/>" + max_date
	static.Customer = c.SearchKeys[1]
	resp.Static = &static
	return
ERR:
	resp.Error = err.Error()
	return
}

func DelSaleRecord(id int) string {
	r := SaleRecord{Id: id}
	// _, err := dborm.Delete(&r)
	r.Deleted = 1
	_, err := dborm.Update(&r, "Deleted")
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
		_, err = dborm.Insert(r)
	} else {
		_, err = dborm.Update(r)
	}

	if err != nil {
		return err.Error()
	}
	return "ok"
}

func SearchCostRecords(c SearchCondition) (resp *Response) {
	table := dborm.QueryTable(&CostRecord{})
	if c.SearchKeys[0] != "" {
		table = table.Filter("date__icontains", c.SearchKeys[0])
	}

	var static CostRecord
	min_date := INIT_MIN_DATE
	var max_date string

	var rows []*CostRecord
	table, resp, err := searchPartial(table, c)
	if err != nil {
		goto ERR
	}

	if resp.Total == 0 {
		return
	}

	_, err = table.All(&rows)
	if err != nil {
		goto ERR
	}

	resp.Rows = rows
	log.Println("SearchCostRecords Total count:", resp.Total, len(rows), rows[0])

	// 统计
	if c.nostatics {
		return
	}
	if resp.Total < SEARCH_MAX_COUNT && resp.Total > len(rows) {
		_, err = table.Limit(SEARCH_MAX_COUNT, 0).All(&rows)
		if err != nil {
			goto ERR
		}
	}
	for _, row := range rows {
		if min_date > row.Date {
			min_date = row.Date
		}
		if max_date < row.Date {
			max_date = row.Date
		}
		static.TeaCost += row.TeaCost
		static.LaborCost += row.LaborCost
		static.Freight += row.Freight
		static.Postage += row.Postage
		static.CartonCost += row.CartonCost
		static.Consumables += row.Consumables
		static.PackingCharges += row.PackingCharges
		static.Total += row.Total
	}
	static.Date = min_date + "-<br/>" + max_date
	resp.Static = &static
	return
ERR:
	resp.Error = err.Error()
	return
}

func DelCostRecord(id int) string {
	r := CostRecord{Id: id}
	// _, err := dborm.Delete(&r)
	r.Deleted = 1
	_, err := dborm.Update(&r, "Deleted")
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
		_, err = dborm.Insert(r)
	} else {
		_, err = dborm.Update(r)
	}

	if err != nil {
		return err.Error()
	}
	return "ok"
}

func SearchStockRecords(c SearchCondition) (resp *Response) {
	table := dborm.QueryTable(&StockRecord{})
	if c.SearchKeys[0] != "" {
		table = table.Filter("date__icontains", c.SearchKeys[0])
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
	// _, err := dborm.Delete(&r)
	r.Deleted = 1
	_, err := dborm.Update(&r, "Deleted")
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
	_, _, err = dborm.ReadOrCreate(&old, "Date")
	if err != nil {
		return err.Error()
	}

	r.Id = old.Id
	_, err = dborm.Update(r)
	if err != nil {
		return err.Error()
	}

	return "ok"
}

func SearchFinanceStaticss(c SearchCondition) (resp *Response) {
	table := dborm.QueryTable(&FinanceStatics{})
	if c.SearchKeys[0] != "" {
		table = table.Filter("date__icontains", c.SearchKeys[0])
	}

	var rows []*FinanceStatics
	table, resp, err := searchPartial(table, c)
	if err != nil {
		goto ERR
	}

	if resp.Total == 0 {
		return
	}

	_, err = table.All(&rows)
	if err != nil {
		goto ERR
	}

	resp.Rows = rows
	log.Println("SearchFinanceStaticss Total count:", resp.Total, len(rows), rows[0])
	return
ERR:
	resp.Error = err.Error()
	return
}

func CalcFinanceByDate(date string) string {
	r := FinanceStatics{Date: date}
	defer log.Println("CalcFinanceByDate", r)
	var resp *Response
	c := SearchCondition{SearchKeys: []string{date, ""}}
	resp = SearchPurchaseRecords(c)
	if resp.Rows != nil {
		for _, row := range resp.Rows.([]*PurchaseRecord) {
			r.Purchase += row.Total
			r.PurchasedStock += row.Weight
		}
	}

	resp = SearchSaleRecords(c)
	if resp.Rows != nil {
		for _, row := range resp.Rows.([]*SaleRecord) {
			r.Sale += row.Total
			r.SaledStock += row.Weight
		}
	}

	resp = SearchCostRecords(c)
	if resp.Rows != nil {
		for _, row := range resp.Rows.([]*CostRecord) {
			r.Cost += row.Total
		}
	}

	predate, err := nextNDate(date, -1)
	if err != nil {
		return err.Error()
	}

	var pre_f = FinanceStatics{Date: predate}
	err = dborm.Read(&pre_f, "Date")
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

func DownloadExcel(table string, c SearchCondition, i18nColumns [][]string) string {
	c.nostatics = true
	c.PageSize = 0
	var resp *Response
	switch table {
	case "采购":
		resp = SearchPurchaseRecords(c)
	case "销售":
		resp = SearchSaleRecords(c)
	case "成本":
		resp = SearchCostRecords(c)
	case "财务":
		resp = SearchFinanceStaticss(c)
	}

	if resp.Rows == nil {
		return "记录为空"
	}

	fileds, zhs := make([]string, 0, len(i18nColumns)), make([]string, 0, len(i18nColumns))
	for _, pair := range i18nColumns {
		fileds = append(fileds, pair[0])
		zhs = append(zhs, pair[1])
	}

	rows := reflectRowsValues(resp.Rows, fileds)
	err := saveExcel(genpath(table, c.SearchKeys...), rows, zhs)
	if err != nil {
		return err.Error()
	}
	return "ok"
}
