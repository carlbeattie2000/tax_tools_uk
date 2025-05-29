package router

import (
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: Global middleware
type (
	PageRenderer func(router *UIRouter, args any) tview.Primitive
	PageHandler  func(path string, args any, next func() PageRenderer) PageRenderer
)

type Route struct {
	path     string
	handlers []PageHandler
}

func newRoute(path string, handlers []PageHandler) *Route {
	return &Route{path, handlers}
}

func nextHandler(path string, args any, handlers []PageHandler) PageRenderer {
	index := 0

	var next func() PageRenderer
	next = func() PageRenderer {
		if index >= len(handlers) {
			return func(_ *UIRouter, _ any) tview.Primitive {
				return nil
			}
		}
		handler := handlers[index]
		index++
		return handler(path, args, next)
	}

	return next()
}

func (route *Route) handleMiddleware(
	router *UIRouter,
	args any,
	globalHandlers []PageHandler,
) tview.Primitive {
	next := nextHandler(route.path, args, append(globalHandlers, route.handlers...))

	return next(router, args)
}

type UIRouter struct {
	routerHistory  Router
	keybindsRouter *KeybindsRouter
	paths          map[string]*Route
	pages          *tview.Pages
	middleware     []PageHandler
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
	handlers ...PageHandler,
) {
	uirouter.paths[path] = newRoute(path, handlers)
}

func (uirouter *UIRouter) RegisterIndex(handlers ...PageHandler) {
	uirouter.RegisterPath("index", handlers...)
	uirouter.Navigate("index", nil)
}

func (uirouter *UIRouter) UseMiddleware(handler PageHandler) {
	uirouter.middleware = append(uirouter.middleware, handler)
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
	renderer := route.handleMiddleware(uirouter, args, uirouter.middleware)
	uirouter.gotoPage(route.path, renderer, true)
}

func (uirouter *UIRouter) Back() {
	uirouter.removeCurrentPage()
	uirouter.routerHistory.back()
	route, ok := uirouter.paths[uirouter.routerHistory.GetCurrentLocation()]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	renderer := route.handleMiddleware(uirouter, nil, uirouter.middleware)
	uirouter.gotoPage(route.path, renderer, true)
}

func (uirouter *UIRouter) Forward() {
	uirouter.removeCurrentPage()
	uirouter.routerHistory.forward()
	route, ok := uirouter.paths[uirouter.routerHistory.GetCurrentLocation()]

	if !ok {
		uirouter.Navigate("not_found", nil)
		return
	}

	renderer := route.handleMiddleware(uirouter, nil, uirouter.middleware)
	uirouter.gotoPage(route.path, renderer, true)
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
