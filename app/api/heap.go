package api

import (
	"encoding/json"
	"net/http"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/algo"
)

// HeapBuild godoc
// @Summary Build a Heap from an int array.
// @Description Builds heap from int array, returns original array and built array
// @Accept  json
// @Produce  json
// @Param data body HeapApiRequest true "Input todo struct"
// @Success 200 {object} HeapApiResponse
// @Router /heap/build [post]
func BuildHeap(w http.ResponseWriter, r *http.Request) {
	Log.Info("Build heap.")

	body, err := GetRequestFromBody(r)
	CheckError(err)

	var heapApiRequest HeapApiRequest

	err = json.Unmarshal(body, &heapApiRequest)
	CheckHttpError(&w, err)

	ArrayRequest := make([]int, len(heapApiRequest.Array))
	copy(ArrayRequest, heapApiRequest.Array)

	Log.Info("HeapApiRequest: ", heapApiRequest)

	var heapApiResponse HeapApiResponse

	var intHeap algo.IntegerHeap = heapApiRequest.Array

	intHeap.BuildMinHeap()

	heapApiResponse.Array = ArrayRequest
	heapApiResponse.BuiltHeap = intHeap

	w.Header().Set("Content-Type", "application/json")
	SetupResponse(&w)
	jsObj, _ := json.Marshal(heapApiResponse)
	w.Write(jsObj)

}

// HeapSort godoc
// @Summary Sort an int array using Heap Sort Algorithm
// @Description Sorts an unsorted int array, returns unordered and ordered array.
// @Accept  json
// @Produce  json
// @Param data body HeapApiRequest true "Input todo struct"
// @Success 200 {object} HeapSortApiResponse
// @Router /heap/sort [post]
func HeapSort(w http.ResponseWriter, r *http.Request) {
	Log.Info("Requesting Heap Sort Algorithm.")

	body, err := GetRequestFromBody(r)
	CheckError(err)

	var heapApiRequest HeapApiRequest

	err = json.Unmarshal(body, &heapApiRequest)
	CheckHttpError(&w, err)

	ArrayRequest := make([]int, len(heapApiRequest.Array))
	copy(ArrayRequest, heapApiRequest.Array)

	Log.Debug("Unordered Array Request: ", heapApiRequest)

	var heapSortApiResponse HeapSortApiResponse

	var intSlice []int = heapApiRequest.Array

	orderedIntSlice := algo.HeapSort(intSlice)

	heapSortApiResponse.UnorderedArray = intSlice
	heapSortApiResponse.OrderedArray = orderedIntSlice

	Log.Debug("Ordered Array Response: ", heapApiRequest)

	w.Header().Set("Content-Type", "application/json")
	SetupResponse(&w)
	jsObj, _ := json.Marshal(heapSortApiResponse)
	w.Write(jsObj)

}

type HeapApiRequest struct {
	Array []int
}

type HeapApiResponse struct {
	Array     []int
	BuiltHeap []int
}

type HeapSortApiResponse struct {
	UnorderedArray []int
	OrderedArray   []int
}
