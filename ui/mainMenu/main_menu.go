package mainmenu

import (
	"tax_calculator/engine/ui/router"

	"github.com/rivo/tview"
)

func GetLayout(router *router.UIRouter) *tview.Flex {
	layout := tview.NewFlex()

	menuForm := tview.NewForm().AddButton("Income Stats", func() {
		router.Navigate("income_stats")
	})

	layout.AddItem(menuForm, 0, 1, true)

	return layout
}
