package profitabilityratios

import "github.com/Natchalit/lib-accounting/src/functions"

// กำไรก่อนดอกเบี้ยและภาษี = กำไรขั้นต้น - ค่าใช้จ่ายในการขายและบริหาร
func ProfitBeforeInterestAndTax(grossProfit, sellingAndAdministrativeExpenses, decimal float64) float64 {
	return functions.Minus(grossProfit, sellingAndAdministrativeExpenses, decimal)
}
