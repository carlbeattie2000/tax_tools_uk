package mainmenu

import (
	"strings"
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NiceName(path string) string {
	caserTitle := cases.Title(language.English)
	caserLower := cases.Lower(language.English)

	return caserTitle.String(caserLower.String(strings.Replace(path, "_", " ", -1)))
}

// Route is added last to router, this way we can generate navigation menu from routers routes
func GetLayout(router *router.UIRouter, _ any) tview.Primitive {
	layout := tview.NewFlex()

	menuForm := tview.NewForm()

	for _, path := range router.GetUserPaths() {
		menuForm.AddButton(NiceName(path), func() {
			router.Navigate(path, nil)
		})
	}

	layout.AddItem(menuForm, 0, 1, true)

	return layout
}
