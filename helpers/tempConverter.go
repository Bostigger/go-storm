package helpers

import "fmt"

func ConvertToCelsius(tempKelvin float64) string {
	convertedVal := tempKelvin - 273.15
	roundedVal := fmt.Sprintf("%.2f", convertedVal)
	return roundedVal
}
