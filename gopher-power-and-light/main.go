package main

import "fmt"

func main() {

	gridLoad := 75.
	activePlants := []int{0, 1}
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please select an option:")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
		}
	case "2":
		var capacity float64

		for _, plantID := range activePlants {
			capacity += plantCapacities[plantID]
		}

		fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
		fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization:", gridLoad/capacity*100)

	default:
		fmt.Println("Unknown option. No action taken.")
	}
}
