package utils

// 计算佣金值
func ComputeCommissionValue(price float64, cv int32) float64 {
	return RoundToOneDecimal(float64(cv) * price / 100)
}
