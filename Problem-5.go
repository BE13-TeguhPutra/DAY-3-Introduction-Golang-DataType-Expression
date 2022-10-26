package main

import (
	"fmt"
	"math"
)

func main() {
	var tinggi float64 = 20
	var radius float64 = 4
	var Luas_tabung float64 = 2 * math.Pi * radius * (radius + tinggi)
	fmt.Println("Luas tabung adalah", Luas_tabung)

}
