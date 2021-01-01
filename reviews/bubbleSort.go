package main

import (
	"fmt"
	"strconv"
)

func main() {
	integers := []int{}
	integers = getIntegers(integers)
	fmt.Println("\nYour integers: ")
	fmt.Println(integers)

	bubbleSort(integers)
	fmt.Println("\nIntegers after sorting: ")
	fmt.Println(integers)
}

func getIntegers(integers []int) []int {
	var userInput string

	for len(integers) < 10 {
		fmt.Println("Please enter an integer.")
		fmt.Scanln(&userInput)

		i, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			integers = append(integers, i)
		}
	}

	return integers
}

func bubbleSort(integers []int) {
	for i := len(integers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if integers[j-1] > integers[j] {
				swap(integers, j, j-1)
			}
		}
	}
}

func swap(integers []int, i int, j int) {
	tmp := integers[j]
	integers[j] = integers[i]
	integers[i] = tmp
}
