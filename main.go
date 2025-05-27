package main

import (
	"net/http"
	notfound "tax_calculator/engine/ui/NotFound"
	incomestats "tax_calculator/engine/ui/incomeStats"
	mainmenu "tax_calculator/engine/ui/mainMenu"
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	app := tview.NewApplication()
	router := router.NewUIRouter(app)

	router.RegisterIndex(mainmenu.GetLayout(router))
	router.RegisterPath("income_stats", incomestats.GetLayout(router))
	router.RegisterPath("not_found", notfound.GetLayout(router))

	router.Navigate("index")

	app.Run()
}
