package main

import "fmt"

func findMax(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	arr := []int{3, 9, 2, 7, 5}
	fmt.Println("Maximum element:", findMax(arr))
}
