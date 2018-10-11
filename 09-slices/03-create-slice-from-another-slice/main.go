package main

import "fmt"

func main() {
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Mumbai", "Bangalore", "Hyderabad", "Hong Kong"}

	asianCities := cities[3:]
	indianCities := asianCities[1:5]

	fmt.Println("cities = ", cities) // cities =  [New York London Chicago Beijing Delhi Mumbai Bangalore Hyderabad Hong Kong]
	fmt.Println("asianCities = ", asianCities) // asianCities =  [Beijing Delhi Mumbai Bangalore Hyderabad Hong Kong]
	fmt.Println("indianCities = ", indianCities) // indianCities =  [Delhi Mumbai Bangalore Hyderabad]
}
