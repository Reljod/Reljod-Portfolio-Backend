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

func BuildHeap(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Build heap.")

	if r.Method != "POST" {
		http.Error(w, "This endpoint only supports POST requests", http.StatusMethodNotAllowed)
		return
	}

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

func HeapSort(w http.ResponseWriter, r *http.Request) {
	logger.Info("Requesting Heap Sort Algorithm.")

	if r.Method != "POST" {
		http.Error(w, "This endpoint only supports POST requests", http.StatusMethodNotAllowed)
		return
	}

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

	var heapSortApiResponse HeapSortApiResponse

	var intSlice []int = heapApiRequest.Array

	orderedIntSlice := algo.HeapSort(intSlice)

	heapSortApiResponse.UnorderedArray = intSlice
	heapSortApiResponse.OrderedArray = orderedIntSlice

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
