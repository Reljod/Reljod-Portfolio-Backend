package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/algo"
	log "github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
)

var logger = log.Logger

// HeapBuild godoc
// @Summary Build a Heap from an int array.
// @Description Builds heap from int array, returns original array and built array
// @Accept  json
// @Produce  json
// @Param data body HeapApiRequest true "Input todo struct"
// @Success 200 {object} HeapApiResponse
// @Router /heap/build [post]
func BuildHeap(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Build heap.")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var heapApiRequest HeapApiRequest

	err = json.Unmarshal(body, &heapApiRequest)
	if err != nil {
		http.Error(w, "Cannot Parse Request Body", http.StatusBadRequest)
		return
	}

	ArrayRequest := make([]int, len(heapApiRequest.Array))
	copy(ArrayRequest, heapApiRequest.Array)

	fmt.Println("HeapApiRequest: ", heapApiRequest)

	var heapApiResponse HeapApiResponse

	var intHeap algo.IntegerHeap = heapApiRequest.Array

	intHeap.BuildMinHeap()

	heapApiResponse.Array = ArrayRequest
	heapApiResponse.BuiltHeap = intHeap

	w.Header().Set("Content-Type", "application/json")
	setupResponse(&w)
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
	logger.Info("Requesting Heap Sort Algorithm.")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var heapApiRequest HeapApiRequest

	err = json.Unmarshal(body, &heapApiRequest)
	if err != nil {
		http.Error(w, "Cannot Parse Request Body", http.StatusBadRequest)
		return
	}

	ArrayRequest := make([]int, len(heapApiRequest.Array))
	copy(ArrayRequest, heapApiRequest.Array)

	logger.Debug("Unordered Array Request: ", heapApiRequest)

	var heapSortApiResponse HeapSortApiResponse

	var intSlice []int = heapApiRequest.Array

	orderedIntSlice := algo.HeapSort(intSlice)

	heapSortApiResponse.UnorderedArray = intSlice
	heapSortApiResponse.OrderedArray = orderedIntSlice

	logger.Debug("Ordered Array Response: ", heapApiRequest)

	w.Header().Set("Content-Type", "application/json")
	setupResponse(&w)
	jsObj, _ := json.Marshal(heapSortApiResponse)
	w.Write(jsObj)

}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
