//go:generate go run -tags generate gen.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"

	"github.com/zserge/lorca"
)

const (
	loginUI = iota
	menuUI
	tableUI

	TIP_HOURS = 24 * 6
	author    = `<a href="mailto:zhoutong.zht@gmail.com?cc=zhoutong.zht@foxmail.com&bcc=zhoutong.zht@gmail.com&subject=请求获取正式版&body=运兴柠檬请求获取正式版授权">作者</a>`
)

var (
	trialday  = "5"
	builddate = "20191130"

	hoursleft  float64
	expireDate time.Time
)

func main() {
	version := flag.Bool("v", false, "version")
	flag.Parse()
	if *version {
		fmt.Printf("%s, %s\n%s built, %s expire\n", runtime.Version(), author, builddate, expireDate.Format("2006/01/02 15:04:05"))
		return
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	var curUI = loginUI

	// NEWUI:
	ui, err := lorca.New("", "", 1400, 900)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close() // 因为循环可能关闭多次

	ui.Bind("loginReady", func() {
		if hoursleft < TIP_HOURS {
			ui.Eval(fmt.Sprintf(`
				$('#expireInfo').removeClass('hide').html('该软件将在%s过期,请联系%s获取正式版');
			`, expireDate.Format("2006/01/02 15:04:05"), author))
		}
	})

	ui.Bind("menuUI", func() {
		curUI = menuUI
	})

	ui.Bind("tableUI", func() {
		curUI = tableUI
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

	// ui.Bind("searchStockRecords", SearchStockRecords)
	// ui.Bind("newOrEditStockRecord", NewOrEditStockRecord)
	// ui.Bind("delStockRecord", DelStockRecord)

	ui.Bind("calcFinanceByPeriod", CalcFinanceByPeriod)
	ui.Bind("searchFinanceStaticss", SearchFinanceStaticss)
	ui.Bind("downloadExcel", DownloadExcel)

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
	case <-ui.Done():
		signal.Stop(sigc)
	}

	log.Println("exiting...")
}

func init() {
	days, _ := strconv.ParseInt(trialday, 10, 0)
	since, _ := time.Parse("20060102", builddate)
	expireDate = since.AddDate(0, 0, int(days))
	hoursleft = expireDate.Sub(time.Now()).Hours()
}
