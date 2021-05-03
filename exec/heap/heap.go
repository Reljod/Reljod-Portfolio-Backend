package main

import (
	"github.com/Reljod/Reljod-Portfolio-Backend/app/algo"
	log "github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
)

var logger = log.Logger

func main() {
	unorderedSlice := []int{5, 4, 3, 2, 1}
	logger.Info("Unordered Slice: ", unorderedSlice)
	orderedSlice := algo.HeapSort(unorderedSlice)
	logger.Info("Ordered Slice: ", orderedSlice)
}
