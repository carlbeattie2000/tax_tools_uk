package app

import (
	"tax_calculator/engine/lib/router"
)

type PageContext struct {
	path     string
	response *router.Response
}

func newPageContext(path string, res *router.Response) *PageContext {
	return &PageContext{path, res}
}
