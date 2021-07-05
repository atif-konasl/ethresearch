package main

import "fmt"

func sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func arrayController() {
	array := [...]float64{7.0, 8.5, 9.1}
	x := sum(&array)  // Note the explicit address-of operator
	fmt.Printf("x=%f", x)
}

