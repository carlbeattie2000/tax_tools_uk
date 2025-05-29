package router

import (
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Route struct {
	path string
	page func(router *UIRouter, args any) tview.Primitive
}

func newRoute(path string, page func(router *UIRouter, args any) tview.Primitive) *Route {
	return &Route{path, page}
}

type UIRouter struct {
	routerHistory  Router
	keybindsRouter *KeybindsRouter
	paths          map[string]*Route
	pages          *tview.Pages
}

func NewUIRouter(app *tview.Application) *UIRouter {
	routerHistory := newRouter(20)

	uiRouter := &UIRouter{
		routerHistory: *routerHistory,
		paths:         make(map[string]*Route),
		pages:         tview.NewPages(),
	}

	app.SetRoot(uiRouter.pages, true).SetFocus(uiRouter.pages)

	uiRouter.keybindsRouter = newKeybindsRouter(app, uiRouter)

	return uiRouter
}

func (uirouter *UIRouter) RegisterPath(
	path string,
	page func(router *UIRouter, args any) tview.Primitive,
) {
	uirouter.paths[path] = newRoute(path, page)
}

func (uirouter *UIRouter) RegisterIndex(page func(router *UIRouter, args any) tview.Primitive) {
	uirouter.RegisterPath("index", page)
	uirouter.Navigate("index", nil)
}

func (uirouter *UIRouter) removeCurrentPage() {
	currentPath := uirouter.routerHistory.GetCurrentLocation()

	if currentPath != "" {
		uirouter.pages.RemovePage(currentPath)
	}
}

func (uirouter *UIRouter) gotoPage(path string, page tview.Primitive, resize bool) {
	uirouter.pages.AddAndSwitchToPage(path, page, resize)
}

func (uirouter *UIRouter) Navigate(path string, args any) {
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	uirouter.removeCurrentPage()
	uirouter.routerHistory.navigate(path)
	uirouter.gotoPage(route.path, route.page(uirouter, args), true)
}

func (uirouter *UIRouter) Back() {
	uirouter.removeCurrentPage()
	uirouter.routerHistory.back()
	route, ok := uirouter.paths[uirouter.routerHistory.GetCurrentLocation()]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	uirouter.gotoPage(route.path, route.page(uirouter, nil), true)
}

func (uirouter *UIRouter) Forward() {
	uirouter.removeCurrentPage()
	uirouter.routerHistory.forward()
	route, ok := uirouter.paths[uirouter.routerHistory.GetCurrentLocation()]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	uirouter.gotoPage(route.path, route.page(uirouter, nil), true)
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

func (uirouter *UIRouter) GetUserPaths() []string {
	nonUserPaths := []string{"index", "not_found"}
	keys := make([]string, 0, len(uirouter.paths))
	for k := range uirouter.paths {
		if slices.Contains(nonUserPaths, k) {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}
