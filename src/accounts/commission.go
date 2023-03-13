package accounts

import "github.com/Natchalit/lib-accounting/src/functions"

// เงินค่าคอมมิชชั้น = ราคาสินค้า * ค่าคอมมิชชั่น
func Commission(price, percent, decimal float64) float64 {
	return functions.DecimalPoint(price*functions.Percent(percent, 2), decimal)
}
