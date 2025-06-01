package router

import (
	"errors"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func helloWorld(req *Request, res *Response, next NextFunc) {
	res.Render(tview.NewFlex().SetTitle("testing_box"))
}

func TestRouterShouldNotStackOverflowWithManyRegisteredRoutes(t *testing.T) {
	router := NewRouter()

	for range 6000 {
		router.Get("/testing", helloWorld)
	}

	router.Get("/", func(req *Request, res *Response, next NextFunc) {
		res.Render(tview.NewBox().SetTitle("test_box"))
	})

	req := NewRequest("/", nil, "")
	res := router.Use(req)

	assert.Equal(t, 200, res.Status)
}

func TestRouterShouldNotStackOverflowWithLargeSyncStack(t *testing.T) {
	router := NewRouter()

	for range 6000 {
		router.Get("/sync_stack", func(req *Request, res *Response, next NextFunc) { next(nil) })
	}

	router.Get("/sync_stack", helloWorld)

	req := NewRequest("/sync_stack", nil, "")
	res := router.Use(req)

	assert.Equal(t, 200, res.Status)
}

func TestShouldBeChainable(t *testing.T) {
	router := NewRouter()
	inner := NewRouter()

	assert.ObjectsAreEqual(router.Get("/", helloWorld), router)
	assert.ObjectsAreEqual(router.UseMiddleware(helloWorld), router)
	assert.ObjectsAreEqual(router.UseRouter(inner), router)
	assert.ObjectsAreEqual(router.UseNamedRouter("/thing", inner), router)
}

func TestShouldSupportAnotherRouter(t *testing.T) {
	router := NewRouter()
	inner := NewRouter()

	inner.Get("/", helloWorld)
	router.UseRouter(inner)

	res := router.Use(NewRequest("/", nil, ""))

	assert.Equal(t, 200, res.Status)
}

func TestShouldAcceptMultipleArguments(t *testing.T) {
	router := NewRouter()

	router.Get("/", func(req *Request, res *Response, next NextFunc) { next(nil) }, helloWorld)

	res := router.Use(NewRequest("/", nil, ""))

	assert.Equal(t, 200, res.Status)
}

func TestShouldNotInvokeSingleErrorFunction(t *testing.T) {
	router := NewRouter()

	router.UseErrorHandler(func(err error, req *Request, res *Response, next NextFunc) {
		panic(err)
	})

	res := router.Use(NewRequest("/", nil, ""))

	// Since this is not a full HTTP server, and the router here doesn't integrate with
	// an HTTP library that would send a 404 by default, the response status remains 200.
	assert.Equal(t, 200, res.Status)
}

func TestShouldInvokeSingleErrorFunction(t *testing.T) {
	router := NewRouter()

	router.UseMiddleware(func(req *Request, res *Response, next NextFunc) {
		next(errors.New("error, oh hello!"))
	})

	router.UseErrorHandler(func(err error, req *Request, res *Response, next NextFunc) {
		res.Status = 500
	})

	res := router.Use(NewRequest("/", nil, ""))

	assert.Equal(t, 500, res.Status)
}

func TestShouldNotInvokeFunctionAboveError(t *testing.T) {
	router := NewRouter()

	router.UseErrorHandler(func(err error, req *Request, res *Response, next NextFunc) {
		res.Status = 500
	})

	router.UseMiddleware(func(req *Request, res *Response, next NextFunc) {
		next(errors.New("error, oh hello!"))
	})

	res := router.Use(NewRequest("/", nil, ""))

	assert.Equal(t, 200, res.Status)
}
