package main

import (
	"net/http"
	"tax_calculator/engine/ui/app"
	mainmenu "tax_calculator/engine/ui/main_menu"
	notfound "tax_calculator/engine/ui/not_found"
	"tax_calculator/engine/ui/router"
	taxcalculator "tax_calculator/engine/ui/tax_calculator"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	app := app.NewApplication()

	app.Get(
		"tax_calculator",
		func(_ string, _ any, _ func() router.PageRenderer) router.PageRenderer {
			return taxcalculator.GetLayout
		},
	)
	app.Get(
		"not_found",
		func(_ string, _ any, _ func() router.PageRenderer) router.PageRenderer {
			return notfound.GetLayout
		},
	)
	app.Get("/", func(_ string, _ any, _ func() router.PageRenderer) router.PageRenderer {
		return mainmenu.GetLayout
	})

	app.Run()
}
