package main

import (
	incomestats "tax_calculator/engine/ui/incomeStats"
	mainmenu "tax_calculator/engine/ui/mainMenu"
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	router := router.NewUIRouter(app)

	router.RegisterIndex(mainmenu.GetLayout(router))
	router.RegisterPath("income_stats", incomestats.GetLayout(router))

	router.Navigate("index")

	app.Run()
}
