package main

import (
	"fmt"
)

func jumpingOnClouds(c []int32) int32 {
	// Write your code here

	var i int
	var convertedClouds []int
	for i < len(c)-3 {
		firstThree := c[i : i+3]
		fmt.Println(firstThree)

		if firstThree[0] != firstThree[2] {
			convertedClouds = append(convertedClouds, 0, 0)
		} else {
			convertedClouds = append(convertedClouds, 0)
		}

		i += 2

		if i == len(c)-2 {
			fmt.Println(c[i:])
			if !(firstThree[0]|firstThree[1] == 1) {
				convertedClouds = append(convertedClouds, 0)
			} else {
				convertedClouds = append(convertedClouds, 0, 0)
			}
		}
	}
	fmt.Println(convertedClouds)

	numClouds := len(convertedClouds) - 1

	return int32(numClouds)
}

func main() {
	numClouds := jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0})
	fmt.Println("Minimum number of clouds to jump on: ", numClouds)
}
