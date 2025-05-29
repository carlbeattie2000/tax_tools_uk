package app

import (
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"
)

type Application struct {
	*tview.Application
	router router.UIRouter
}

func NewApplication() *Application {
	tviewApp := tview.NewApplication()
	return &Application{tviewApp, *router.NewUIRouter(tviewApp)}
}

func (app *Application) Get(path string, handlers ...router.PageHandler) {
	if path == "/" || path == "index" {
		app.router.RegisterIndex(handlers...)
		return
	}

	app.router.RegisterPath(path, handlers...)
}
