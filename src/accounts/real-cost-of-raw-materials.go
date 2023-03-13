package accounts

import "github.com/Natchalit/lib-accounting/src/functions"

// ทุนจริงของวัตถุดิบต่อชิ้น = ต้นทุนรวมทั้งหมด / จำนวนชิ้นที่ได้
func RealCostOfRawMaterials(totalCostPrice, pieces, decimal float64) float64 {
	return functions.Divine(totalCostPrice, pieces, decimal)
}
