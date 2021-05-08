package main

import "fmt"

func main() {
	var intArr []int = []int{2, 4, 6, 8, 10, 1, 3, 5, 7, 9}
	var toFind = 3

	isFound, foundAt := findAnItem(&intArr, toFind)
	if isFound {
		fmt.Println(fmt.Sprint(`Value `, toFind, ` found at position `, foundAt))
	} else {
		fmt.Println(fmt.Sprint(`Value `, toFind, ` not found`))
	}
}

func findAnItem(value *[]int, toFind int) (isFound bool, foundAt int) {

	foundAt = -1
	isFound = false

	for index, value := range *value {
		if value == toFind {
			isFound = true
			foundAt = index
			break
		}
	}
	return
}
