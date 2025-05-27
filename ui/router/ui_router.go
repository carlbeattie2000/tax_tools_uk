package router

import (
	"tax_calculator/engine/internal/logger"

	"github.com/rivo/tview"
)

type Route struct {
	path string
	page *tview.Flex
}

func newRoute(path string, page *tview.Flex) *Route {
	return &Route{path, page}
}

type UIRouter struct {
	routerHistory Router
	app           *tview.Application
	paths         map[string]*Route
}

func NewUIRouter(app *tview.Application) *UIRouter {
	routerHistory := newRouter(20)

	uiRouter := &UIRouter{
		routerHistory: *routerHistory,
		app:           app,
		paths:         make(map[string]*Route),
	}

	return uiRouter
}

func (uirouter *UIRouter) RegisterIndex(page *tview.Flex) {
	uirouter.paths["index"] = newRoute("index", page)
}

func (uirouter *UIRouter) RegisterPath(path string, page *tview.Flex) {
	uirouter.paths[path] = newRoute(path, page)
}

func (uirouter *UIRouter) Navigate(path string) {
	logger.GetLogger().Printf("navigating to path: %s\n", path)
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found")
		return
	}

	uirouter.routerHistory.Navigate(path)
	uirouter.app.SetRoot(route.page, true)
	logger.GetLogger().Println("navigated")
}

func (uirouter *UIRouter) Back() {
	uirouter.routerHistory.Back()
	path := uirouter.routerHistory.currentPageNode.page
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found")
		return
	}

	uirouter.app.SetRoot(route.page, true)
}

func (uirouter *UIRouter) Forward() {
	uirouter.routerHistory.Forward()
	path := uirouter.routerHistory.currentPageNode.page
	route, ok := uirouter.paths[path]

	if !ok {
		uirouter.Navigate("not_found")
		return
	}

	uirouter.app.SetRoot(route.page, true)
}
