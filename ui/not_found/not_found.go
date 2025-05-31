package notfound

import (
	"tax_calculator/engine/ui/app"

	"github.com/rivo/tview"
)

func GetLayout(app *app.Application) tview.Primitive {
	layout := tview.NewFlex()

	notFoundForm := tview.NewForm().
		AddTextView("Error", "Page Not Found", 20, 2, true, false).
		AddButton("home", func() {
			app.Fetch("/", nil, "")
		})
	layout.AddItem(notFoundForm, 0, 1, true)

	return layout
}
