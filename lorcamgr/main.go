//go:generate go run -tags generate gen.go

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

const (
	loginUI = iota
	menuUI
	tableUI
)

func main() {
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
	defer ui.Close()

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
	case <-ui.Done():
		// log.Println("curUI", curUI, tableUI)
		if curUI == tableUI {
			signal.Stop(sigc)
			goto NEWUI
		}
	}

	log.Println("exiting...")
}
