package notfound

import (
	"tax_calculator/engine/lib/app"

	"github.com/rivo/tview"
)

func GetLayout(app *app.Application) tview.Primitive {
	layout := tview.NewFlex()

	notFoundForm := tview.NewForm().
		AddTextView("Error", "Page Not Found", 20, 2, true, false).
		AddButton("home", func() {
			app.Fetch("/views/", nil, "")
		})
	layout.AddItem(notFoundForm, 0, 1, true)

	return layout
}
