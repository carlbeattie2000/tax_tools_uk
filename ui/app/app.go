package app

import (
	"tax_calculator/engine/lib/router"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Application struct {
	*tview.Application
	*router.Router
	*KeybindsRouter
}

func NewApplication() *Application {
	tviewApp := tview.NewApplication()
	app := &Application{tviewApp, router.NewRouter(tviewApp), nil}
	app.KeybindsRouter = newKeybindsRouter(app)
	return app
}

func (app *Application) Fetch(path string, params map[string]string, query string) {
	req := router.NewRequest(path, params, query)
	app.Use(req)
}

func (app *Application) Start() {
	app.Router.RegisterDefaultHandlers()
	app.Fetch("/views/", nil, "")
	app.Run()
}

// TODO: Support new routing features like middleware
type KeybindsRouter struct {
	app      *Application
	keybinds map[tcell.Key]string
}

func newKeybindsRouter(app *Application) *KeybindsRouter {
	kr := KeybindsRouter{app, map[tcell.Key]string{}}

	kr.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if location, ok := kr.keybinds[event.Key()]; ok {
			kr.app.Fetch(location, nil, "")
			return nil
		}

		return event
	})

	return &kr
}

func (keyboardrouter *KeybindsRouter) RegisterKey(key tcell.Key, location string) {
	keyboardrouter.keybinds[key] = location
}
