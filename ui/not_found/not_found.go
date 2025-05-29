package notfound

import (
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"
)

func GetLayout(router *router.UIRouter, _ any) tview.Primitive {
	layout := tview.NewFlex()

	notFoundForm := tview.NewForm().
		AddTextView("Error", "Page Not Found", 20, 2, true, false).
		AddButton("home", func() {
			router.Navigate("index", nil)
		})
	layout.AddItem(notFoundForm, 0, 1, true)

	return layout
}
