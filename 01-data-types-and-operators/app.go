package main

import "fmt"

func main() {
	capacity := 0.
	// var capacity float64
	gridLoad := 75.
	// var gridLoad float64 = 75.
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	fmt.Println(plantCapacities)

	for i := 0; i < len(plantCapacities); i++ {
		capacity += plantCapacities[i]
	}

	utilization := (gridLoad / capacity) * 100

	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", utilization)

}
