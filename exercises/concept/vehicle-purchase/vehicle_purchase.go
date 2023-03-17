package purchase

import (
	"math"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	requires := [2]string{"car", "truck"}
	for _, val := range requires {
		if val == kind {
			return true
		}
	}
	return false
	// return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	m := [][]float64{
		{3.0, 0.8},
		{10.0, 0.7},
		{math.Inf(1), 0.5},
	}
	for _, pair := range m {
		if age < pair[0] {
			return pair[1] * originalPrice
		}
	}
	panic("unreachable")
	/*
		// 第一种解法
		var rate float64
		if age < 3 {
			rate = 0.8
		} else if age < 10 {
			rate = 0.7
		} else {
			rate = 0.5
		}
		return rate * originalPrice
	*/
}
