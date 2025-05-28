package router

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Route struct {
	path string
	page func(router *UIRouter, args any) *tview.Flex
}

func newRoute(path string, page func(router *UIRouter, args any) *tview.Flex) *Route {
	return &Route{path, page}
}

type UIRouter struct {
	routerHistory  Router
	keybindsRouter *KeybindsRouter
	app            *tview.Application
	paths          map[string]*Route
}

func NewUIRouter(app *tview.Application) *UIRouter {
	routerHistory := newRouter(20)

	uiRouter := &UIRouter{
		routerHistory: *routerHistory,
		app:           app,
		paths:         make(map[string]*Route),
	}

	uiRouter.keybindsRouter = newKeybindsRouter(app, uiRouter)

	return uiRouter
}

func (uirouter *UIRouter) RegisterIndex(page func(router *UIRouter, args any) *tview.Flex) {
	uirouter.paths["index"] = newRoute("index", page)
}

func (uirouter *UIRouter) RegisterPath(
	path string,
	page func(router *UIRouter, args any) *tview.Flex,
) {
	uirouter.paths[path] = newRoute(path, page)
}

func (uirouter *UIRouter) Navigate(path string, args any) {
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found", "well, this is awkward")
		return
	}

	uirouter.routerHistory.navigate(path)
	uirouter.app.SetRoot(route.page(uirouter, args), true)
}

func (uirouter *UIRouter) Back() {
	uirouter.routerHistory.back()
	path := uirouter.routerHistory.location.location
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	uirouter.app.SetRoot(route.page(uirouter, nil), true)
}

func (uirouter *UIRouter) Forward() {
	uirouter.routerHistory.forward()
	path := uirouter.routerHistory.location.location
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	uirouter.app.SetRoot(route.page(uirouter, nil), true)
}

func (uirouter *UIRouter) ListenerSubscribe(
	subscriberFunc func(update *RouterUpdate),
) {
	uirouter.routerHistory.listener.subscribe(subscriberFunc)
}

func (uirouter *UIRouter) RegisterKeybind(key tcell.Key, location string) {
	if _, ok := uirouter.paths[location]; ok {
		uirouter.keybindsRouter.registerKey(key, location)
	}
}
