package viewrouter

import (
	"tax_calculator/engine/lib/app"
	r "tax_calculator/engine/lib/router"
	mainmenu "tax_calculator/engine/ui/main_menu"
	taxcalculator "tax_calculator/engine/ui/tax_calculator"
)

func ViewRouter(app *app.Application) *r.Router {
	router := app.NewRouter()

	router.Get("/", func(req *r.Request, res *r.Response, next r.NextFunc) {
		res.Render(mainmenu.GetLayout(app))
	})
	router.Get("/tax_calculator", func(req *r.Request, res *r.Response, next r.NextFunc) {
		res.Render(taxcalculator.GetLayout(app))
	})

	return router
}
