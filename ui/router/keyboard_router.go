package router

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TODO: Support new routing features like middleware
type KeybindsRouter struct {
	app      *tview.Application
	router   *UIRouter
	keybinds map[tcell.Key]string
}

func newKeybindsRouter(app *tview.Application, router *UIRouter) *KeybindsRouter {
	kr := KeybindsRouter{app, router, map[tcell.Key]string{}}

	kr.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if location, ok := kr.keybinds[event.Key()]; ok {
			kr.router.Navigate(location, nil)
			return nil
		}

		return event
	})

	return &kr
}

func (keyboardrouter *KeybindsRouter) registerKey(key tcell.Key, location string) {
	keyboardrouter.keybinds[key] = location
}
