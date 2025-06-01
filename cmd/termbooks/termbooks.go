package main

import (
	"net/http"
	"tax_calculator/engine/internal/logger"
	"tax_calculator/engine/lib/app"
	"tax_calculator/engine/lib/router"
	notfound "tax_calculator/engine/ui/not_found"
	viewrouter "tax_calculator/engine/ui/view_router"

	_ "net/http/pprof"

	"github.com/gdamore/tcell/v2"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	tuiApp := app.NewApplication()
	viewRouter := viewrouter.ViewRouter(tuiApp)

	tuiApp.UseNamedRouter("/views", viewRouter)

	tuiApp.Middleware(func(req *router.Request, res *router.Response, next router.NextFunc) {
		res.Render(notfound.GetLayout(tuiApp))
	})

	tuiApp.ErrorHandler(
		func(err error, req *router.Request, res *router.Response, next router.NextFunc) {
			logger.GetLogger().Println(err)
		},
	)

	tuiApp.RegisterKeybind(tcell.KeyF1, func(app *app.Application) {
		app.Back()
	})
	tuiApp.RegisterKeybind(tcell.KeyF2, func(app *app.Application) {
		app.Forward()
	})

	tuiApp.RunWithInitialPath("/views/")
}
