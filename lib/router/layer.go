package router

import (
	"errors"
)

type (
	NextFunc           func(error)
	RequestHandlerFunc func(req *Request, res *Response, next NextFunc)
	ErrorFunc          func(err error, req *Request, res *Response, next NextFunc)
)

type Layer struct {
	path  string
	keys  []string
	route *Route

	middlewareHandler RequestHandlerFunc
	errorHandler      ErrorFunc
}

func NewRouteLayer(path string, keys []string) *Layer {
	return &Layer{path: path, keys: keys}
}

func NewMiddlewareLayer(path string, keys []string, handler RequestHandlerFunc) *Layer {
	return &Layer{
		path:              path,
		keys:              keys,
		middlewareHandler: handler,
	}
}

func NewErrorMiddlewareLayer(path string, keys []string, handler ErrorFunc) *Layer {
	return &Layer{path: path, keys: keys, errorHandler: handler}
}

func (layer *Layer) HandleError(err error, req *Request, res *Response, next NextFunc) {
	fn := layer.errorHandler

	if fn != nil {
		fn(err, req, res, next)
		return
	}

	next(err)
}

func (layer *Layer) HandleRequest(req *Request, res *Response, next NextFunc) {
	if layer.route != nil {
		layer.route.dispatch(req, res, next)
		return
	}

	if layer.middlewareHandler != nil {
		layer.middlewareHandler(req, res, next)
		return
	}

	next(errors.New("route"))
}

// TODO: Need to do actual matching
func (layer *Layer) Match(path string) bool {
	if layer.path == "*" {
		return true
	}
	return layer.path == path
}

type LayerStack []*Layer
