package router

import (
	"errors"

	"github.com/rivo/tview"
)

type RouterHistoryAction string

const (
	POP  RouterHistoryAction = "POP"
	PUSH RouterHistoryAction = "PUSH"
)

type RouterHistoryNode struct {
	context  *RequestContext
	previous *RouterHistoryNode
	next     *RouterHistoryNode
}

func newRouterHistoryNode(ctx *RequestContext) *RouterHistoryNode {
	return &RouterHistoryNode{context: ctx}
}

// RouterHistory represents the routers history
type RouterHistory struct {
	head     *RouterHistoryNode
	location *RouterHistoryNode
	size     int
	maxSize  int
}

func newHistoryRouter(maxSize int) *RouterHistory {
	return &RouterHistory{maxSize: maxSize}
}

func (rh *RouterHistory) trimHead(delta int) {
	if delta <= 0 {
		return
	}
	if rh.size == 1 {
		rh.head = nil
		rh.location = nil
		rh.size = 0

		return
	}
	for rh.head != nil && delta > 0 {
		rh.head = rh.head.next
		rh.head.previous = nil
		rh.size--
		delta--
	}
}

func (rh *RouterHistory) nodeLengthToEndFromNode(node *RouterHistoryNode) int {
	var count int
	for node.next != nil {
		count++
		node = node.next
	}
	return count
}

func (rh *RouterHistory) addNode(node *RouterHistoryNode) {
	newSize := rh.size + 1
	if newSize > rh.maxSize {
		rh.trimHead(1)
	}

	if rh.head == nil {
		rh.head = node
		rh.location = node
	} else {
		sizeToRemove := rh.nodeLengthToEndFromNode(rh.location)
		rh.size -= sizeToRemove
		node.previous = rh.location
		rh.location.next = node
		rh.location = node
	}

	rh.size++
}

func (rh *RouterHistory) navigate(ctx *RequestContext) {
	historyNode := newRouterHistoryNode(ctx)
	rh.addNode(historyNode)
}

func (rh *RouterHistory) forward() {
	if rh.head == nil || rh.location.next == nil {
		return
	}

	rh.location = rh.location.next
}

func (rh *RouterHistory) back() {
	if rh.head == nil || rh.location.previous == nil {
		return
	}

	rh.location = rh.location.previous
}

type Router struct {
	history *RouterHistory
	pages   *tview.Pages
	stack   LayerStack
}

func NewRouter(app *tview.Application) *Router {
	r := &Router{history: newHistoryRouter(20), pages: tview.NewPages(), stack: LayerStack{}}
	if app != nil {
		app.SetRoot(r.pages, true).SetFocus(r.pages)
	}
	return r
}

func (router *Router) handle(req *Request, res *Response) {
	var idx int
	var sync int
	var stack LayerStack = router.stack
	var next NextFunc

	next = func(err error) {
		var layerError error

		if err != nil {
			if err.Error() == "route" {
				layerError = nil
			} else {
				layerError = err

				if layerError.Error() == "router" {
					next(layerError)
					return
				}
			}
		}

		if idx >= len(stack) {
			next(layerError)
			return
		}

		if sync++; sync > 100 {
			next(errors.New("sync exceeded"))
			return
		}

		if req.path == "" {
			next(layerError)
			return
		}

		var layer *Layer
		var match bool
		var route *Route

		for match != true && idx < len(stack) {
			layer = stack[idx]
			match = layer.Match(req.path)
			route = layer.route
			idx++

			if match != true {
				continue
			}

			if route == nil {
				continue
			}
		}

		if match != true {
			next(layerError)
			return
		}

		if layerError != nil {
			layer.HandleError(layerError, req, res, next)
		} else {
			layer.HandleRequest(req, res, next)
		}

		sync = 0
	}

	next(nil)
}

func (router *Router) Route(path string) *Route {
	route := newRoute(path)
	layer := NewRouteLayer(path, []string{})
	layer.route = route
	router.stack = append(router.stack, layer)
	return route
}

func (router *Router) Get(path string, handlers ...RequestHandlerFunc) *Route {
	route := newRoute(path)
	route.Get(handlers...)
	layer := NewRouteLayer(path, []string{})
	layer.route = route
	router.stack = append(router.stack, layer)
	return route
}

func (router *Router) UseMiddleware(handlers ...RequestHandlerFunc) {
	for _, handler := range handlers {
		layer := NewMiddlewareLayer("*", nil, handler)
		router.stack = append(router.stack, layer)
	}
}

func (router *Router) UseNamedMiddleware(path string, handlers ...RequestHandlerFunc) {}

func (router *Router) UseRouter(newRouter *Router) {
	for _, layer := range newRouter.stack {
		router.stack = append(router.stack, layer)
	}
}

func (router *Router) UseNamedRouter(path string, newRouter *Router) {}

func (router *Router) UseErrorHandler(handlers ...ErrorFunc) {
	for _, handler := range handlers {
		router.stack = append(router.stack, NewErrorMiddlewareLayer("*", []string{}, handler))
	}
}

func (router *Router) Use(req *Request) {
	res := newResponse(router)
	requestContext := newRequestContext(req)
	router.history.navigate(requestContext)

	router.handle(req, res)
}

func (router *Router) Back() {
	router.history.back()
	req := router.CurrentHistoryLocationContext().Request
	res := newResponse(router)
	router.handle(req, res)
}

func (router *Router) Forward() {
	router.history.forward()
	req := router.CurrentHistoryLocationContext().Request
	res := newResponse(router)
	router.handle(req, res)
}

func (router *Router) CurrentHistoryLocationContext() *RequestContext {
	return router.history.location.context
}

func (router *Router) GetPaths() []string {
	var paths []string
	for _, layer := range router.stack {
		if layer.route != nil && layer.route.path != "/" {
			paths = append(paths, layer.route.path)
		}
	}
	return paths
}
