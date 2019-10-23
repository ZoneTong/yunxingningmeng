package main

import (
	"strings"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

const (
	SEARCH_DATE = iota
	SEARCH_PROVIDER
	SEARCH_CUSTOMER
)

type SearchCondition struct {
	SearchArea int    `json:"searchArea"`
	SearchKey  string `json:"searchKey"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	SortName   string `json:"sortName"`
	SortOrder  string `json:"sortOrder"` // desc asc
}

type Response struct {
	// TotalNotFiltered int64       `json:"totalNotFiltered"`
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
	Error string
}

func nextNDate(date string, n int) (nd string, err error) {
	t, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return "", err
	}
	return t.AddDate(0, 0, n).Format(DATE_FORMAT), nil
}

func structFieldName2DBCloumnName(s string) string {
	var n strings.Builder
	n.WriteRune(unicode.ToLower([]rune(s)[0]))
	// log.Println("structFieldName2DBCloumnName", n.String())
	for _, r := range s[1:] {
		if unicode.IsUpper(r) {
			n.WriteRune('_')
			n.WriteRune(unicode.ToLower(r))
		} else {
			n.WriteRune(r)
		}
	}
	// log.Println("structFieldName2DBCloumnName2", n.String())
	return n.String()
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
