package router

type Route struct {
	path  string
	stack LayerStack
}

func newRoute(path string) *Route {
	return &Route{path, make(LayerStack, 0)}
}

func (route *Route) dispatch(req *Request, res *Response, done NextFunc) {
	var idx int
	var sync int
	var stack LayerStack = route.stack

	if len(stack) == 0 {
		done(nil)
		return
	}

	var next NextFunc
	next = func(err error) {
		if err != nil {
			switch err.Error() {
			case "route":
				done(nil)
				return
			case "router":
				done(err)
				return
			default:
				break
			}
		}

		if idx >= len(stack) {
			done(err)
			return
		}

		if sync++; sync > 100 {
			done(err)
			return
		}

		var layer *Layer = stack[idx]
		idx++

		if err != nil {
			layer.HandleError(err, req, res, next)
		} else {
			layer.HandleRequest(req, res, next)
		}

		sync = 0
	}

	next(nil)
}

func (route *Route) Get(handlers ...RequestHandlerFunc) {
	for _, handler := range handlers {
		layer := NewMiddlewareLayer(route.path, []string{}, handler)
		route.stack = append(route.stack, layer)
	}
}
