// Imagine you are developing software for a weather station.
// This station records temperatures throughout the day.

// Write a programa in Go that process a list of temperatures
// and provides the highest, lowest, and average temperature.

// Steps:
// Initialized a slice of integers representing temperatures
// at different times of the day.
//
// Iterate through the slice to find the maximum and minimum temperatures.
//
// Calculate the average temperature of the day
//
// Print ou the maximum, minimum, and average temperatures.

package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{15.20, 11.30, 18.6, 22.1, 23.9}

	sort.Float64s(temperatures)

	fmt.Println(temperatures)
	fmt.Printf("Lowest temperature: %.2f\n", temperatures[0])
	fmt.Printf("Highest temperature: %.2f\n", temperatures[len(temperatures)-1])

	var sum float64

	for _, temperature := range temperatures {
		sum += temperature
	}

	average := sum / float64(len(temperatures))
	fmt.Printf("Average temperature: %.2f\n", average)
}
