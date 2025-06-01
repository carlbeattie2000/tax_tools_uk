package app

import "github.com/gdamore/tcell/v2"

type KeybindHandler func(app *Application)

type Keybinds struct {
	app      *Application
	keybinds map[tcell.Key]KeybindHandler
}

func newKeybinds(app *Application) *Keybinds {
	kr := Keybinds{app, map[tcell.Key]KeybindHandler{}}

	kr.app.tui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if handler, ok := kr.keybinds[event.Key()]; ok {
			handler(app)
			return nil
		}

		return event
	})

	return &kr
}

func (keyboardrouter *Keybinds) registerKey(key tcell.Key, handler KeybindHandler) {
	keyboardrouter.keybinds[key] = handler
}
