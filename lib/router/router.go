package router

import (
	"tax_calculator/engine/internal/logger"
)

type Router struct {
	stack LayerStack
}

func NewRouter() *Router {
	r := &Router{LayerStack{}}
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
			return
		}

		if sync++; sync > 100 {
			go next(err)
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

		sync = 0

		if layerError != nil {
			layer.HandleError(layerError, req, res, next)
		} else {
			layer.HandleRequest(req, res, next)
		}
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

func (router *Router) Get(path string, handlers ...RequestHandlerFunc) *Router {
	route := newRoute(path)
	route.Get(handlers...)
	layer := NewRouteLayer(path, []string{})
	layer.route = route
	router.stack = append(router.stack, layer)
	return router
}

func (router *Router) UseMiddleware(handlers ...RequestHandlerFunc) *Router {
	for _, handler := range handlers {
		layer := NewMiddlewareLayer("*", nil, handler)
		router.stack = append(router.stack, layer)
	}
	return router
}

func (router *Router) UseNamedMiddleware(path string, handlers ...RequestHandlerFunc) {}

func (router *Router) UseRouter(newRouter *Router) *Router {
	for _, layer := range newRouter.stack {
		router.stack = append(router.stack, layer)
	}
	return router
}

func (router *Router) UseNamedRouter(path string, newRouter *Router) *Router {
	for _, layer := range newRouter.stack {
		layer.path = path + layer.path
		if layer.route != nil {
			layer.route.path = path + layer.route.path
		}
		router.stack = append(router.stack, layer)
	}
	return router
}

func (router *Router) UseErrorHandler(handlers ...ErrorFunc) {
	for _, handler := range handlers {
		router.stack = append(router.stack, NewErrorMiddlewareLayer("*", []string{}, handler))
	}
}

func (router *Router) Use(req *Request) *Response {
	res := NewResponse()

	router.handle(req, res)

	return res
}

func handle404(req *Request, res *Response, next NextFunc) {
	res.Status = 404
}

func (router *Router) register404Handler() {
	if len(router.stack) == 0 {
		router.UseMiddleware(handle404)
		return
	}

	var found404Layer *Layer

	for i := len(router.stack) - 1; i >= 0; i-- {
		if router.stack[i].route != nil {
			break
		}

		if router.stack[i].middlewareHandler != nil {
			found404Layer = router.stack[i]
			break
		}
	}

	if found404Layer == nil {
		router.UseMiddleware(handle404)
	}
}

func handleError(err error, req *Request, res *Response, next NextFunc) {
	res.Status = 500

	logger.GetLogger().Println(err, res.Status)
}

func (router *Router) registerErrorHandler() {
	if len(router.stack) == 0 {
		router.UseErrorHandler(handleError)
		return
	}

	lastLayer := router.stack[len(router.stack)-1]

	if lastLayer.errorHandler == nil {
		router.UseErrorHandler(handleError)
	}
}

func (router *Router) RegisterDefaultHandlers() {
	router.register404Handler()
	router.registerErrorHandler()
}
