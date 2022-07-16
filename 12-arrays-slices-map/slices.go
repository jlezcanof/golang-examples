package main

import (
	"fmt"
)

func slices() {

	beers := make([]string, 3)
	fmt.Println("empty slice:", beers)

	beers[0] = "estrella"
	beers[1] = "bucker"
	beers[2] = "mahou"
	fmt.Println("beers:", beers)
	fmt.Println("beers[2]:", beers[2])

	beers = append(beers, "coronita", "san miguel")
	fmt.Println("beers:", beers)

	other := beers[:3]

	fmt.Println("other beers:", other)

	other[0] = "bucker 0.0"
	fmt.Println("other beers:", other)
	fmt.Println("beers:", beers)
}
