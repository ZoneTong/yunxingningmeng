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

	// Create and bind Go object to the UI
	// r := &PurchaseRecord{Id: 3, Date: time.Now().Format("20060102"), Weight: 3, Price: 4}
	ui.Bind("verifyPassword", verifyPassword)
	ui.Bind("searchPurchaseRecords", SearchPurchaseRecords)
	ui.Bind("newOrEditPurchaseRecord", NewOrEditPurchaseRecord)
	ui.Bind("delPurchaseRecord", DelPurchaseRecord)

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
