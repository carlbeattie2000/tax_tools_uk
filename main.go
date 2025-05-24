package main

import (
	incomestats "tax_calculator/engine/ui/incomeStats"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	taxStatsView := incomestats.GetLayout()

	root := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(taxStatsView, 0, 1, true)

	app.SetRoot(root, true).Run()
}
