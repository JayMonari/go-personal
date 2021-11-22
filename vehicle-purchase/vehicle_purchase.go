package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	if n := strings.Compare(option1, option2); n > 0 {
		return fmt.Sprint(option2, " is clearly the better choice.")
	}
	return fmt.Sprint(option1, " is clearly the better choice.")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 10 {
		return originalPrice * 0.5
	} else if age >= 3 {
		return originalPrice * 0.7
	}
	return originalPrice * 0.8
}
