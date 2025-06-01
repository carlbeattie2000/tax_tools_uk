package main

import (
	"net/http"
	"tax_calculator/engine/internal/logger"
	"tax_calculator/engine/lib/app"
	"tax_calculator/engine/lib/router"
	notfound "tax_calculator/engine/ui/not_found"
	viewrouter "tax_calculator/engine/ui/view_router"

	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	app := app.NewApplication()
	viewRouter := viewrouter.ViewRouter(app)

	app.UseNamedRouter("/views", viewRouter)

	app.Middleware(func(req *router.Request, res *router.Response, next router.NextFunc) {
		res.Render(notfound.GetLayout(app))
	})

	app.ErrorHandler(
		func(err error, req *router.Request, res *router.Response, next router.NextFunc) {
			logger.GetLogger().Println(err)
		},
	)

	app.RunWithInitialPath("/views/")
}
