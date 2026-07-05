// Package cars calculates multiple values related to car production.
package cars

// costPerCar is cost per car in batches less than 10.
const costPerCar = 10000

// costPerTenCar is cost for 10 cars produced in batch.
const costPerTenCar = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (float64(successRate) / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var total uint
	total += (uint(carsCount) / 10) * costPerTenCar
	total += (uint(carsCount) % 10) * costPerCar

	return total
}
