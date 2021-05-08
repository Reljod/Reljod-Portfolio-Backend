package main

import (
	"fmt"
	"strconv"
)

/*
There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

Example


There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .

Function Description

Complete the sockMerchant function in the editor below.

sockMerchant has the following parameter(s):

int n: the number of socks in the pile
int ar[n]: the colors of each sock
Returns

int: the number of pairs
Input Format

The first line contains an integer , the number of socks represented in .
The second line contains  space-separated integers, , the colors of the socks in the pile.
*/

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var numPairs map[string]int32 = make(map[string]int32)
	for _, sock := range ar {
		fmt.Println(sock)
		numSock := strconv.Itoa(int(sock))

		_, ok := numPairs[numSock]
		if !ok {
			numPairs[numSock] = 1
		} else {
			numPairs[numSock] += 1
		}
	}

	fmt.Println(numPairs)

	var numPairsCount int32
	for _, value := range numPairs {
		value = value / 2
		numPairsCount += value
	}

	return numPairsCount
}

func main() {
	numPairsCount := sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	fmt.Println("Number of Pairs: ", numPairsCount)
}
