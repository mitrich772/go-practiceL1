package main

import (
	"fmt"
	"math"
)

func groupTemp(temps []float64) map[int][]float64 {
	var roundedTemp int
	groupedTemps := make(map[int][]float64)
	for _, currentTemp := range temps {
		roundedTemp = int(math.Floor(currentTemp/10)) * 10
		groupedTemps[roundedTemp] = append(groupedTemps[roundedTemp], currentTemp) // Если такого нету то вернет в append nil а тот создаст из 1го элемента уже
	}
	return groupedTemps
}
func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Print(groupTemp(temps))
}
