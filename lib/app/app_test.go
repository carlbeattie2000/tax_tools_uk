package app

import (
	"errors"
	"tax_calculator/engine/lib/router"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAppHasDefault404Handler(t *testing.T) {
	app := NewApplication()
	done := make(chan struct{})

	go func() {
		app.Run()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)

	res := app.Fetch("/", nil, "")

	app.tui.QueueUpdate(func() {
		app.Stop()
	})

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("App did not stop in time")
	}

	assert.Equal(t, 404, res.Status)
}

func TestAppHasDefaultErrorHandler(t *testing.T) {
	app := NewApplication()
	done := make(chan struct{})

	go func() {
		app.Get("/", func(req *router.Request, res *router.Response, next router.NextFunc) {
			next(errors.New("oh a error"))
		})

		app.Run()
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)

	res := app.Fetch("/", nil, "")

	app.tui.QueueUpdate(func() {
		app.Stop()
	})

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("App did not stop in time")
	}

	assert.Equal(t, 500, res.Status)
}
