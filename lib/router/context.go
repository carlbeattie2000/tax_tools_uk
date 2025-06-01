package router

import (
	"github.com/rivo/tview"
)

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

type ResponseHeader map[string]string

type Response struct {
	View     tview.Primitive
	Status   int
	Req      *Request
	Redirect string
	headers  ResponseHeader
}

func NewResponse() *Response {
	return &Response{Status: 200}
}

func (res *Response) SendStatus(status int) *Response {
	res.Status = status
	return res
}

func (res *Response) Render(view tview.Primitive) *Response {
	res.View = view
	return res
}

func (res *Response) SendRedirect(path string) *Response {
	res.Redirect = path
	return res
}

func (res *Response) Set(name string, value string) *Response {
	res.headers[name] = value
	return res
}

func (res *Response) Get(name string) string {
	value, ok := res.headers[name]

	if !ok {
		return ""
	}

	return value
}

func (res *Response) Type(headerType string) *Response {
	res.headers["Content-Type"] = headerType
	return res
}
