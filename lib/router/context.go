package router

import "github.com/rivo/tview"

type Request struct {
	path   string
	params map[string]string
	query  string // TODO: Add support for query strings
}

func NewRequest(path string, params map[string]string, query string) *Request {
	return &Request{path, params, query}
}

func (req *Request) GetPath() string {
	return req.path
}

type RequestContext struct {
	*Request
	view tview.Primitive
}

func newRequestContext(request *Request) *RequestContext {
	return &RequestContext{request, nil}
}

func (rc *RequestContext) SetView(page tview.Primitive) {
	rc.view = page
}

type Response struct {
	*Router
	status int
}

func newResponse(router *Router) *Response {
	return &Response{router, 200}
}

func (res *Response) Render(page tview.Primitive) {
	currentLocationContext := res.CurrentHistoryLocationContext()
	if currentLocationContext.view != nil {
		res.pages.AddAndSwitchToPage(currentLocationContext.path, currentLocationContext.view, true)
		return
	}
	res.pages.AddAndSwitchToPage(currentLocationContext.path, page, true)
	currentLocationContext.SetView(page)
}
