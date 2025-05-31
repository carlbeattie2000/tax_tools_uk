package viewrouter

import (
	r "tax_calculator/engine/lib/router"
	"tax_calculator/engine/ui/app"
	mainmenu "tax_calculator/engine/ui/main_menu"
	taxcalculator "tax_calculator/engine/ui/tax_calculator"
)

func ViewRouter(app *app.Application) *r.Router {
	router := r.NewRouter(nil)

	router.Get("/", func(req *r.Request, res *r.Response, next r.NextFunc) {
		res.Render(mainmenu.GetLayout(app))
	})
	router.Get("/tax_calculator", func(req *r.Request, res *r.Response, next r.NextFunc) {
		res.Render(taxcalculator.GetLayout(app))
	})

	return router
}
