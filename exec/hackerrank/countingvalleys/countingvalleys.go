package main

import (
	"fmt"
)

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var valleyCount int32
	var altitude int

	for _, curPath := range path {
		var curAltitude = altitude
		if string(curPath) == "U" {
			altitude += 1
		} else {
			altitude -= 1
		}

		if curAltitude == 0 && altitude == -1 {
			valleyCount += 1
		}
	}

	return valleyCount
}

func main() {
	fmt.Println("Counting Valleys")

	// var steps int = 8
	// var path []string = []string{"D", "D", "U", "U", "U", "U", "D", "D"}
	// return number of Valleys.

	valleyCount := countingValleys(8, "DDUUUUDDUUUUDDUDUDUDUDUDDDDDDUU")
	fmt.Println("Valley counts", valleyCount)

}
