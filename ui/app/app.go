package app

import (
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"
)

type Application struct {
	*tview.Application
	router.UIRouter
}

func NewApplication() *Application {
	tviewApp := tview.NewApplication()
	return &Application{tviewApp, *router.NewUIRouter(tviewApp)}
}

func (app *Application) Get(path string, handlers ...router.PageHandler) {
	if path == "/" || path == "index" {
		app.RegisterIndex(handlers...)
		return
	}

	app.RegisterPath(path, handlers...)
}
