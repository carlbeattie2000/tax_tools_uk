package incomestats

import (
	"fmt"
	"strconv"
	utilmath32 "tax_calculator/engine/internal/utils/util_math32"
	taxliability "tax_calculator/engine/internal/valueobjects/tax_liability"
	"tax_calculator/engine/ui/router"

	"github.com/chewxy/math32"
	"github.com/rivo/tview"
)

type TaxStats struct {
	income                float32
	taxPaid               float32
	taxPercentageOfIncome float32
}

func SetTaxStat(income float32, scan *TaxStats) {
	scan.income = income
	scan.taxPaid = taxliability.CalculateTaxLiability(income)
	scan.taxPercentageOfIncome = utilmath32.PercentageOfTotal(scan.taxPaid, scan.income)
}

// TODO: Show income take home pay increase
// TODO: Show more stats around tax increase
// TODO: Split into tools
// TODO: Tool that can incrementally calculate taxes, and what to put side per pay -- Like if you are a freelancer, and you get paid x, then what does your yearly total now come too, and what do you have to pay out of what you just earned

// Account -- Account number -- 4 digit code && Name -- Both can be used to find accounts
// SQLite with option to setup connections to other SQL Databases
// Configurable tax year
// Deployable as a service so you can communicate via API (Ideal when you accept multiple small payments)

func GetLayout(appRouter *router.UIRouter) *tview.Flex {
	income1 := float32(0)
	income2 := float32(0)
	stat1 := TaxStats{}
	stat2 := TaxStats{}

	taxResultViewIncome := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income 1:")
	taxResultViewTaxPaid := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income 1 Tax Paid:")
	taxResultViewTaxPaidPercent := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income 1 Tax Paid Percent:")

	taxResultViewIncome2 := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income 2:")
	taxResultViewTaxPaid2 := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income 2 Tax Paid:")
	taxResultViewTaxPaidPercent2 := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Income 2 Tax Paid Percent:")

	largerTaxIncreasePercentage := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetLabel("Higher Income Tax Increase From Lower Percent:")

	form := tview.NewForm()

	form.AddInputField("Income 1: £", "", 20, nil, func(text string) {
		if s, err := strconv.ParseFloat(text, 32); err == nil {
			income1 = float32(s)
		}
	})
	form.AddInputField("Income 2: £", "", 20, nil, func(text string) {
		if s, err := strconv.ParseFloat(text, 32); err == nil {
			income2 = float32(s)
		}
	})
	form.AddButton("submit", func() {
		SetTaxStat(income1, &stat1)
		SetTaxStat(income2, &stat2)

		taxResultViewIncome.SetText(fmt.Sprintf("£%f", stat1.income))
		taxResultViewTaxPaid.SetText(fmt.Sprintf("£%f", stat1.taxPaid))
		taxResultViewTaxPaidPercent.SetText(fmt.Sprintf("%f", stat1.taxPercentageOfIncome) + "%")

		taxResultViewIncome2.SetText(fmt.Sprintf("£%f", stat2.income))
		taxResultViewTaxPaid2.SetText(fmt.Sprintf("£%f", stat2.taxPaid))
		taxResultViewTaxPaidPercent2.SetText(fmt.Sprintf("%f", stat2.taxPercentageOfIncome) + "%")

		higher := math32.Max(stat1.taxPercentageOfIncome, stat2.taxPercentageOfIncome)
		lower := math32.Min(stat1.taxPercentageOfIncome, stat2.taxPercentageOfIncome)
		increasePercentage := utilmath32.IncreasePercentage(lower, higher)

		largerTaxIncreasePercentage.SetText(fmt.Sprintf("%f", increasePercentage) + "%")
	}).SetFieldBackgroundColor(000).SetButtonBackgroundColor(000)

	form.AddButton("back", func() {
		appRouter.Navigate("well")
	})

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(taxResultViewIncome, 3, 0, false).
		AddItem(taxResultViewTaxPaid, 3, 0, false).
		AddItem(taxResultViewTaxPaidPercent, 3, 0, false).
		AddItem(taxResultViewIncome2, 3, 0, false).
		AddItem(taxResultViewTaxPaid2, 3, 0, false).
		AddItem(taxResultViewTaxPaidPercent2, 3, 0, false).
		AddItem(largerTaxIncreasePercentage, 3, 0, false)

	return layout
}
