package accounts

import "github.com/Natchalit/lib-accounting/src/functions"

// รวมรายรับ
func Income(nums []float64, decimal float64) float64 {
	return functions.SumTotal(nums, decimal)
}
