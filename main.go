package main

import (
	"net/http"
	"tax_calculator/engine/internal/logger"
	"tax_calculator/engine/lib/router"
	"tax_calculator/engine/ui/app"
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

	app.UseRouter(viewRouter)

	app.UseMiddleware(func(req *router.Request, res *router.Response, next router.NextFunc) {
		res.Render(notfound.GetLayout(app))
	})

	app.UseErrorHandler(
		func(err error, req *router.Request, res *router.Response, next router.NextFunc) {
			logger.GetLogger().Println(err)
		},
	)

	app.Start()
}
