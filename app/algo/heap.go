package algo

import (
	"fmt"

	log "github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
)

var logger = log.Logger

type IntegerHeap []int

func (iheap IntegerHeap) Len() int       { return len(iheap) }
func (iheap *IntegerHeap) Swap(i, j int) { (*iheap)[i], (*iheap)[j] = (*iheap)[j], (*iheap)[i] }
func (iheap *IntegerHeap) Push(heapintf int) {
	*iheap = append(*iheap, heapintf)
}

func (iheap *IntegerHeap) GetParent(index int) (int, int) {
	parentIndex := (index - 1) / 2
	return (*iheap)[parentIndex], parentIndex
}

func (iheap *IntegerHeap) GetLeftChild(parentIndex int) (int, int) {
	childLeftIndex := 2*parentIndex + 1
	return (*iheap)[childLeftIndex], childLeftIndex
}

func (iheap *IntegerHeap) GetRightChild(parentIndex int) (int, int) {
	childRightIndex := 2*parentIndex + 2
	return (*iheap)[childRightIndex], childRightIndex
}

func (iheap *IntegerHeap) GetMinChild(parentIndex int) (int, int) {

	if parentIndex*2+2 >= iheap.Len() {
		childLeft, childLeftIndex := iheap.GetLeftChild(parentIndex)
		return childLeft, childLeftIndex
	} else {
		childLeft, childLeftIndex := iheap.GetLeftChild(parentIndex)
		childRight, childRightIndex := iheap.GetRightChild(parentIndex)

		if childLeft < childRight {
			return childLeft, childLeftIndex
		} else {
			return childRight, childRightIndex
		}
	}
}

func (iheap *IntegerHeap) HasChild(index int) bool {
	if index*2+1 < iheap.Len() {
		return true
	} else {
		return false
	}
}

func (iheap *IntegerHeap) BuildMinHeap() {
	heapLen := iheap.Len()

	for currentIndex := heapLen - 1; currentIndex > -1; {

		currentElement := (*iheap)[currentIndex]

		logger.WithFields(map[string]interface{}{
			"CurrentIndex":   currentIndex,
			"CurrentElement": currentElement,
		}).Debug("Iterating through elements.")

		icurrentIndex := currentIndex
		for iheap.HasChild(icurrentIndex) {
			icurrentElement := (*iheap)[icurrentIndex]
			child, childIx := iheap.GetMinChild(icurrentIndex)

			logger.WithFields(map[string]interface{}{
				"ICurrentIndex":   icurrentIndex,
				"ICurrentElement": icurrentElement,
				"MinChild":        child,
				"IMinChild":       childIx,
			}).Debug("Iterating through heap's index with Child.")

			if child < icurrentElement {
				iheap.Swap(childIx, icurrentIndex)

				logger.WithFields(map[string]interface{}{"Child": child, "ICurrentElement": icurrentElement}).Debug("Swapped Child and Current Element.")
			}
			icurrentIndex = childIx
			logger.WithFields(map[string]interface{}{"NewHeap": iheap}).Debug("Looking at heap that was transformed")
		}

		if currentIndex == 0 {
			break
		}

		currentIndex = currentIndex - 1
	}
}

func (iheap *IntegerHeap) BuildMaxHeap() {
	heapLen := iheap.Len()

	for currentIndex := heapLen - 1; currentIndex > -1; {

		currentElement := (*iheap)[currentIndex]

		logger.WithFields(map[string]interface{}{
			"CurrentIndex":   currentIndex,
			"CurrentElement": currentElement,
		}).Debug("Iterating through elements.")

		icurrentIndex := currentIndex
		for iheap.HasChild(icurrentIndex) {
			icurrentElement := (*iheap)[icurrentIndex]
			child, childIx := iheap.GetMinChild(icurrentIndex)

			logger.WithFields(map[string]interface{}{
				"ICurrentIndex":   icurrentIndex,
				"ICurrentElement": icurrentElement,
				"MinChild":        child,
				"IMinChild":       childIx,
			}).Debug("Iterating through heap's index with Child.")

			if child > icurrentElement {
				iheap.Swap(childIx, icurrentIndex)

				logger.WithFields(map[string]interface{}{"Child": child, "ICurrentElement": icurrentElement}).Debug("Swapped Child and Current Element.")
			}
			icurrentIndex = childIx
			logger.WithFields(map[string]interface{}{"NewHeap": iheap}).Debug("Looking at heap that was transformed")
		}

		if currentIndex == 0 {
			break
		}

		currentIndex = currentIndex - 1
	}
}

func (iheap *IntegerHeap) SwapChild(parentIndex int) {
	if parentIndex*2+2 >= iheap.Len() {
		return
	} else {
		childLeft, childLeftIndex := iheap.GetLeftChild(parentIndex)
		childRight, childRightIndex := iheap.GetRightChild(parentIndex)

		if childRight < childLeft {
			iheap.Swap(childRightIndex, childLeftIndex)
		}
	}
}

func (iheap *IntegerHeap) IsCurentHeapHasChildren(index int) bool {
	return index*2+2 < iheap.Len()
}

func (iheap *IntegerHeap) ReorderHeapAfterPush() {
	heapLen := len(*iheap)
	lastIndex := heapLen - 1
	lastHeapElement := (*iheap)[lastIndex]

	var currentIndex int = lastIndex
	var currentHeapElement int = lastHeapElement

	parentElement, parentIndex := iheap.GetParent(currentIndex)

	for currentHeapElement < parentElement && currentIndex != parentIndex {
		logger.Debug("Parent Element: ", parentElement, "Parent index: ", parentIndex)
		logger.Debug("Current index", currentIndex)
		logger.Debug(*iheap)
		iheap.Swap(currentIndex, parentIndex)
		logger.Debug(*iheap)

		currentIndex = parentIndex
		parentElement, parentIndex = iheap.GetParent(currentIndex)

		logger.Debug("New Parent Element: ", parentElement, "New Parent index: ", parentIndex)
		logger.Debug("New Current Element:", currentHeapElement, "New Current index", currentIndex)
	}
}

func (iheap *IntegerHeap) ReorderHeapAfterPop() {
	currentIndex := 0
	for iheap.HasChild(currentIndex) {

		currentElement := (*iheap)[currentIndex]
		child, childIx := iheap.GetMinChild(currentIndex)
		if child < currentElement {
			iheap.Swap(childIx, currentIndex)
		}

		currentIndex = childIx
	}

}

func HeapSort(unorderedSlice []int) []int {
	var intHeap IntegerHeap
	intHeap = append(intHeap, unorderedSlice...)

	logger.Debug("Original Heap: ", intHeap)
	intHeap.BuildMinHeap()
	logger.Debug("Min Heap: ", intHeap)

	var orderedList []int

	for intHeap.Len() > 0 {
		minVal := intHeap.PopMinValue()
		logger.Debug("Popped Min Value Slice: ", intHeap)
		logger.Debug("Min Value of Slice: ", minVal)

		orderedList = append(orderedList, minVal)

		intHeap.ReorderHeapAfterPop()
		logger.Debug("Reordered Popped Slice: ", intHeap)
	}

	logger.Debug("Original List: ", unorderedSlice)
	logger.Debug("Ordered List: ", orderedList)

	return orderedList
}

func (iheap *IntegerHeap) PopMinValue() int {

	minValue := (*iheap)[0]

	iheap.Swap(0, iheap.Len()-1)
	*iheap = (*iheap)[:iheap.Len()-1]

	return minValue
}

func ImplementMinHeap() {
	var intHeap IntegerHeap = IntegerHeap{10, 9, 8, 7, 10, 9, 8, 1, 3, 5}

	fmt.Println("Original Slice: ", intHeap)
	intHeap.BuildMinHeap()
	fmt.Println("Built Heap Slice: ", intHeap)

	// fmt.Println("Push element: 1")
	// intHeap.Push(1)
	// fmt.Println("Heap after pushing: ", intHeap)
	// intHeap.ReorderHeapAfterPush()
	// fmt.Println("Heap after reordering after push: ", intHeap)

	// minValue := intHeap.PopMinValue()
	// fmt.Println("Popped Min Value of Heap: ", minValue)
	// fmt.Println("Heap after popping Min Value", intHeap)
	// intHeap.BuildMinHeap()
	// fmt.Println("Heap after rebuilding", intHeap)
}
