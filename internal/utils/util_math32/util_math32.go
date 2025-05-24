package utilmath32

import "github.com/chewxy/math32"

func RoundDown(value float32, places float32) float32 {
	places = math32.Floor(places)
	divisor := math32.Pow(10, places)

	return math32.Floor(value*divisor) / divisor
}

func RoundUp(value float32, places float32) float32 {
	places = math32.Floor(places)
	divisor := math32.Pow(10, places)

	return math32.Ceil(value*divisor) / divisor
}

func Round(value float32, places float32) float32 {
	places = math32.Floor(places)
	divisor := math32.Pow(10, places)

	return math32.Round(value*divisor) / divisor
}

// IncreasePercentage() rounds to the nearest 2 decimal places
func IncreasePercentage(startingValue float32, finalValue float32) float32 {
	return Round(((finalValue-startingValue)/startingValue)*100, 2)
}

func PercentageOfTotal(part float32, total float32) float32 {
	return Round((part/total)*100, 2)
}
