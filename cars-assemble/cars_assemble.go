// Package cars is used to calculate the production rate of the assembly line
package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {

	rate := 0.0

	if speed == 0 {
		rate = 0.0
	} else if speed >= 1 && speed <= 4 {
		rate = 1.0
	} else if speed >= 5 && speed <= 8 {
		rate = 0.9
	} else {
		rate = 0.77
	}

	return rate
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed*221) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	productionRatePerHour := CalculateProductionRatePerHour(speed)
	return int(productionRatePerHour / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	productionRatePerHour := CalculateProductionRatePerHour(speed)
	if productionRatePerHour > limit {
		return limit
	}

	return productionRatePerHour
}
