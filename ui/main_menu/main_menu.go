package mainmenu

import (
	"strings"
	"tax_calculator/engine/ui/app"

	"github.com/rivo/tview"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NiceName(path string) string {
	caserTitle := cases.Title(language.English)
	caserLower := cases.Lower(language.English)

	return caserTitle.String(
		caserLower.String(strings.Replace(strings.Replace(path, "/", "", -1), "_", " ", -1)),
	)
}

// Route is added last to router, this way we can generate navigation menu from routers routes
func GetLayout(app *app.Application) tview.Primitive {
	layout := tview.NewFlex()

	menuForm := tview.NewForm()

	for _, path := range app.GetPaths() {
		menuForm.AddButton(NiceName(path), func() {
			app.Fetch(path, nil, "")
		})
	}

	layout.AddItem(menuForm, 0, 1, true)

	return layout
}
