package main

import "fmt"

func main() {
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.

	fmt.Println("1) Generate power plant report")
	fmt.Println("2) Grid load report")
	fmt.Println("choose an option")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		plantCapacityReport(plantCapacities...)
	case "2":
		plantGridReport(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("invalid option")

	}
}

func plantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func plantGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}
	fmt.Printf("%-20s%.0f\n", "Capacity : ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load : ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization : ", gridLoad/capacity*100)
}
