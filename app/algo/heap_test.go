package algo

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func HeapLenShouldBeCorrect(intHeap IntegerHeap, expectedLen int) string {
	intHeapLen := intHeap.Len()
	if intHeapLen != expectedLen {
		return fmt.Sprintf("Len() of %v should be %v, got %v.", intHeap, expectedLen, intHeapLen)
	}
	return ""
}

func TestHeapLenShouldBeCorrect(t *testing.T) {
	var testOutput string
	if testOutput = HeapLenShouldBeCorrect(IntegerHeap{1, 2, 3, 4, 5}, 5); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapLenShouldBeCorrect(IntegerHeap{1}, 1); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapLenShouldBeCorrect(IntegerHeap{}, 0); testOutput != "" {
		t.Errorf(testOutput)
	}
}

func HeapSwapShouldBeCorrect(intHeap IntegerHeap, expectedIntHeap IntegerHeap, i int, j int) string {
	intHeap.Swap(i, j)
	if !cmp.Equal(intHeap, expectedIntHeap) {
		return fmt.Sprintf("Integer Heap %v should be equal to expected swapped Int Heap %v", intHeap, expectedIntHeap)
	}
	return ""
}

func TestHeapSwapShouldBeCorrect(t *testing.T) {
	var testOutput string
	if testOutput = HeapSwapShouldBeCorrect(
		IntegerHeap{1, 2, 3},
		IntegerHeap{2, 1, 3},
		0, 1); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapSwapShouldBeCorrect(
		IntegerHeap{5, 4, 3, 2, 10},
		IntegerHeap{5, 10, 3, 2, 4},
		1, 4); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapSwapShouldBeCorrect(
		IntegerHeap{1, 5, 9, 10},
		IntegerHeap{1, 5, 10, 9},
		2, 3); testOutput != "" {
		t.Errorf(testOutput)
	}
}

func HeapPushShouldBeCorrect(intHeap IntegerHeap, expectedIntHeap IntegerHeap, pushedElement int) string {
	intHeap.Push(pushedElement)
	if !cmp.Equal(intHeap, expectedIntHeap) {
		return fmt.Sprintf("Integer Heap %v should be equal to expected Pusheed Int Heap %v", intHeap, expectedIntHeap)
	}
	return ""
}

func TestHeapPushShouldBeCorrect(t *testing.T) {
	var testOutput string
	if testOutput = HeapPushShouldBeCorrect(
		IntegerHeap{1, 2, 3},
		IntegerHeap{1, 2, 3, 4},
		4); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapPushShouldBeCorrect(
		IntegerHeap{5, 4, 3, 2, 10},
		IntegerHeap{5, 4, 3, 2, 10, 10},
		10); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapPushShouldBeCorrect(
		IntegerHeap{1, 5, 9, 10},
		IntegerHeap{1, 5, 9, 10, 2},
		2); testOutput != "" {
		t.Errorf(testOutput)
	}
}

func ReorderHeapAfterPushShouldBeCorrect(intHeap IntegerHeap, expectedIntHeap IntegerHeap) string {
	intHeap.ReorderHeapAfterPush()
	if !cmp.Equal(intHeap, expectedIntHeap) {
		return fmt.Sprintf("Integer Heap %v should be equal to expected Pusheed Int Heap %v", intHeap, expectedIntHeap)
	}
	return ""
}

func TestReorderHeapAfterPushShouldBeCorrect(t *testing.T) {
	var testOutput string
	if testOutput = ReorderHeapAfterPushShouldBeCorrect(
		IntegerHeap{1, 3, 5, 7, 9, 11, 13, 2},
		IntegerHeap{1, 2, 5, 3, 9, 11, 13, 7}); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = ReorderHeapAfterPushShouldBeCorrect(
		IntegerHeap{2, 5, 4, 7, 6, 1},
		IntegerHeap{1, 5, 2, 7, 6, 4}); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = ReorderHeapAfterPushShouldBeCorrect(
		IntegerHeap{4, 9, 8, 10, 11, 9, 10, 11, 12, 12, 13, 15, 3},
		IntegerHeap{3, 9, 4, 10, 11, 8, 10, 11, 12, 12, 13, 15, 9}); testOutput != "" {
		t.Errorf(testOutput)
	}
}

func BuildMinHeapShouldBeCorrect(intArray IntegerHeap, expectedIntHeap IntegerHeap) string {
	intArray.BuildMinHeap()
	if !cmp.Equal(intArray, expectedIntHeap) {
		return fmt.Sprintf("Integer Array %v should be equal to expected Built Int Heap %v", intArray, expectedIntHeap)
	}
	return ""
}

func TestBuildMinHeapShouldBeCorrect(t *testing.T) {
	var testOutput string
	if testOutput = BuildMinHeapShouldBeCorrect(
		IntegerHeap{5, 4, 3, 2, 1, 6, 7},
		IntegerHeap{1, 2, 3, 5, 4, 6, 7}); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = BuildMinHeapShouldBeCorrect(
		IntegerHeap{3, 2, 1},
		IntegerHeap{1, 2, 3}); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = BuildMinHeapShouldBeCorrect(
		IntegerHeap{10, 8, 9, 6, 7, 1, 2, 3},
		IntegerHeap{1, 3, 2, 6, 7, 9, 10, 8}); testOutput != "" {
		t.Errorf(testOutput)
	}
}

func HeapSortShouldBeCorrect(unorderedIntSlice []int, expectedOrderedIntSlice []int) string {
	orderedSlice := HeapSort(unorderedIntSlice)
	if !cmp.Equal(orderedSlice, expectedOrderedIntSlice) {
		return fmt.Sprintf("Integer Unordered Array %v ordered to %v should be equal to expected Ordered Array %v",
			unorderedIntSlice,
			orderedSlice,
			expectedOrderedIntSlice)
	}
	return ""
}

func TestHeapSort(t *testing.T) {
	var testOutput string
	if testOutput = HeapSortShouldBeCorrect(
		[]int{5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5}); testOutput != "" {
		t.Errorf(testOutput)
	}
	if testOutput = HeapSortShouldBeCorrect(
		[]int{8, 3, 1, 2, 10, 5, 4, 6, 7, 9},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}); testOutput != "" {
		t.Errorf(testOutput)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unorderedIntSlice := []int{8, 3, 1, 2, 10, 5, 4, 6, 7, 9}
		_ = HeapSort(unorderedIntSlice)
	}
}
