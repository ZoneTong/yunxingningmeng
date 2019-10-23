//go:generate go run -tags generate gen.go

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/zserge/lorca"
)

const (
	loginUI = iota
	menuUI
	tableUI
)

var (
	TRIAL_DAY  = "8"
	BUILD_TIME = "20191010"
)

func main() {
	hoursleft := expireDate().Sub(time.Now()).Hours()
	if hoursleft < 0 {
		log.Fatalln("试用期已过,请联系作者zhoutong.zht@gmail.com ")
		return
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	var curUI = loginUI

NEWUI:
	ui, err := lorca.New("", "", 400, 250)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close() // 因为循环可能关闭多次

	ui.Bind("loginReady", func() {
		if hoursleft < 24*3 {
			ui.Eval(fmt.Sprintf(`
				$('#expireInfo').removeClass('hide').html('该软件将在%s过期,请联系作者zhoutong.zht@gmail.com获取正式版');
			`, expireDate().Format("2006/01/02 15:04:05")))
		}
	})

	ui.Bind("menuUI", func() {
		curUI = menuUI
		bd, _ := ui.Bounds()
		bd.Width = 400
		bd.Height = 400
		ui.SetBounds(bd)
	})

	ui.Bind("tableUI", func() {
		curUI = tableUI
		bd, _ := ui.Bounds()
		bd.Width = 1350
		bd.Height = 800
		ui.SetBounds(bd)
	})

	ui.Bind("tableUI2", func() {
		curUI = tableUI
		bd, _ := ui.Bounds()
		bd.Width = 1120
		bd.Height = 800
		ui.SetBounds(bd)
	})

	ui.Bind("verifyPassword", verifyPassword)
	ui.Bind("searchPurchaseRecords", SearchPurchaseRecords)
	ui.Bind("newOrEditPurchaseRecord", NewOrEditPurchaseRecord)
	ui.Bind("delPurchaseRecord", DelPurchaseRecord)
	ui.Bind("searchSaleRecords", SearchSaleRecords)
	ui.Bind("newOrEditSaleRecord", NewOrEditSaleRecord)
	ui.Bind("delSaleRecord", DelSaleRecord)

	ui.Bind("searchCostRecords", SearchCostRecords)
	ui.Bind("newOrEditCostRecord", NewOrEditCostRecord)
	ui.Bind("delCostRecord", DelCostRecord)

	ui.Bind("searchStockRecords", SearchStockRecords)
	ui.Bind("newOrEditStockRecord", NewOrEditStockRecord)
	ui.Bind("delStockRecord", DelStockRecord)

	ui.Bind("calcFinanceByPeriod", CalcFinanceByPeriod)
	ui.Bind("searchFinanceStaticss", SearchFinanceStaticss)
	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	pageurl := fmt.Sprintf("http://%s", ln.Addr())
	if curUI == tableUI {
		pageurl += "/menu.html"
	}
	ui.Load(pageurl)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	select {
	case <-sigc:
		ui.Close()

	case <-ui.Done():
		// log.Println("curUI", curUI, tableUI)
		if curUI == tableUI {
			signal.Stop(sigc)
			ui.Close()
			goto NEWUI
		}
	}

	log.Println("exiting...")
}

func NewUI(curUI int, addr string) {

}

func expireDate() time.Time {
	days, _ := strconv.ParseInt(TRIAL_DAY, 10, 0)
	since, _ := time.Parse("20060102", BUILD_TIME)
	return since.AddDate(0, 0, int(days))
}
