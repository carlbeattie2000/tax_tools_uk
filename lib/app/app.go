package app

import (
	"tax_calculator/engine/lib/router"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Application struct {
	tui      *tview.Application
	router   *router.Router
	pages    *tview.Pages
	history  *History
	keybinds *Keybinds
}

func NewApplication() *Application {
	tui := tview.NewApplication()
	pages := tview.NewPages()
	tui.SetRoot(pages, true).SetFocus(pages)
	router := router.NewRouter()
	history := newHistory(15)
	app := &Application{tui, router, pages, history, nil}
	keybinds := newKeybinds(app)
	app.keybinds = keybinds
	return app
}

func (app *Application) setPage(page tview.Primitive) {
	app.pages.RemovePage("view")
	app.pages.AddAndSwitchToPage("view", page, true)
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
	return router.NewRouter()
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
	res := app.router.Use(router.NewRequest(path, params, query))

	app.history.navigate(newPageContext(path, res))

	if res.Status >= 200 && res.Status < 300 && res.View != nil {
		app.setPage(res.View)
		return res
	}

	if res.Redirect != "" {
		app.Fetch(res.Redirect, nil, "")
		return res
	}

	return res
}

func (app *Application) Run() {
	app.router.RegisterDefaultHandlers()
	err := app.tui.Run()
	if err != nil {
		panic(app)
	}
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

func (app *Application) Stop() {
	app.tui.Stop()
}

func (app *Application) Back() {
	app.history.back()
	ctx := app.history.location.context
	if ctx == nil {
		return
	}
	if ctx.response.View != nil {
		app.setPage(ctx.response.View)
	}
}

func (app *Application) Forward() {
	app.history.forward()
	ctx := app.history.location.context
	if ctx == nil {
		return
	}
	if ctx.response.View != nil {
		app.setPage(ctx.response.View)
	}
}

func (app *Application) RegisterKeybind(key tcell.Key, handler KeybindHandler) {
	app.keybinds.registerKey(key, handler)
}
