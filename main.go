package main

import (
	"net/http"
	mainmenu "tax_calculator/engine/ui/main_menu"
	notfound "tax_calculator/engine/ui/not_found"
	"tax_calculator/engine/ui/router"
	taxcalculator "tax_calculator/engine/ui/tax_calculator"

	"github.com/rivo/tview"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	app := tview.NewApplication()
	router := router.NewUIRouter(app)

	router.RegisterPath("tax_calculator", taxcalculator.GetLayout)
	router.RegisterPath("not_found", notfound.GetLayout)

	router.RegisterIndex(mainmenu.GetLayout)

	app.Run()
}
