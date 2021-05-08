package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
)

var Log = log.Logger

func CheckError(err error) {
	if err != nil {
		Log.Fatal(err)
	}
}

func CheckHttpError(w *http.ResponseWriter, err error) {
	if err != nil {
		http.Error(*w, "Cannot Parse Request Body", http.StatusBadRequest)
		return
	}
}

func GetRequestFromBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func WriteJsonToResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	SetupResponse(&w)
	jsObj, _ := json.Marshal(body)
	w.Write(jsObj)
}

func SetupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
