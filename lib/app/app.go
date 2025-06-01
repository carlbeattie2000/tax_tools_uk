package app

import (
	"tax_calculator/engine/lib/router"
	"time"

	"github.com/rivo/tview"
)

type Application struct {
	tui    *tview.Application
	router *router.Router
}

func NewApplication() *Application {
	tui := tview.NewApplication()
	return &Application{tui, router.NewRouter(tui)}
}

func (app *Application) Get(path string, handlers ...router.RequestHandlerFunc) *Application {
	app.router.Get(path, handlers...)
	return app
}

func (app *Application) Middleware(handlers ...router.RequestHandlerFunc) *Application {
	app.router.UseMiddleware(handlers...)
	return app
}

func (app *Application) ErrorHandler(handlers ...router.ErrorFunc) *Application {
	app.router.UseErrorHandler(handlers...)
	return app
}

func (app *Application) Route(path string) *router.Route {
	return app.router.Route(path)
}

func (app *Application) NewRouter() *router.Router {
	return router.NewRouter(nil)
}

func (app *Application) UseRouter(router *router.Router) *Application {
	app.router.UseRouter(router)
	return app
}

func (app *Application) UseNamedRouter(path string, router *router.Router) *Application {
	app.router.UseNamedRouter(path, router)
	return app
}

func (app *Application) Fetch(
	path string,
	params map[string]string,
	query string,
) *router.Response {
	return app.router.Use(router.NewRequest(path, params, query))
}

func (app *Application) Run() {
	app.router.RegisterDefaultHandlers()
	err := app.tui.Run()
	panic(err)
}

func (app *Application) RunWithInitialPath(path string) {
	go func() {
		time.Sleep(10 * time.Millisecond)
		app.tui.QueueUpdateDraw(func() {
			app.router.RegisterDefaultHandlers()
			app.Fetch(path, nil, "")
		})
	}()
	err := app.tui.Run()
	panic(err)
}

// // TODO: Support new routing features like middleware
// type KeybindsRouter struct {
// 	app      *Application
// 	keybinds map[tcell.Key]string
// }
//
// func newKeybindsRouter(app *Application) *KeybindsRouter {
// 	kr := KeybindsRouter{app, map[tcell.Key]string{}}
//
// 	kr.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 		if location, ok := kr.keybinds[event.Key()]; ok {
// 			kr.app.Fetch(location, nil, "")
// 			return nil
// 		}
//
// 		return event
// 	})
//
// 	return &kr
// }
//
// func (keyboardrouter *KeybindsRouter) RegisterKey(key tcell.Key, location string) {
// 	keyboardrouter.keybinds[key] = location
// }
