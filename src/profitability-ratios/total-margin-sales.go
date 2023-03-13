package profitabilityratios

import "github.com/Natchalit/lib-accounting/src/functions"

// กําไรส่วนเกินรวมยอดขาย = ยอดขาย - ต้นทุนผันแปรรวม ยอดขาย
func TotalMarginSales(sales, totalVariableCostsSales, decimal float64) float64 {
	return functions.Minus(sales, totalVariableCostsSales, decimal)
}
