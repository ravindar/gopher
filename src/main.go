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
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
		}
	case "2":
		capacity := 0.
		for _, plantID := range activePlants {
			capacity += plantCapacities[plantID]
		}
		fmt.Printf("%-20s%.0f\n", "Capacity : ", capacity)
		fmt.Printf("%-20s%.0f\n", "Load : ", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization : ", gridLoad/capacity*100)
	default:
		fmt.Println("invalid option")

	}
}
