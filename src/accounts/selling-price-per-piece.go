package accounts

import "github.com/Natchalit/lib-accounting/src/functions"

// ราคาขายต่อชิ้น = ต้นทุนรวมต่อชิ้น + ((กำไร(%) * ต้นทุนรวมต่อชิ้น) / 100)
func SellingPricePerPiece(totalCostPrice, profit, decimal float64) float64 {
	profitPerPrice := (profit * totalCostPrice) / 100
	return functions.Sum(totalCostPrice, profitPerPrice, decimal)
}
