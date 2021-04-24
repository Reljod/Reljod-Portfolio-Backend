package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world!")

	router := mux.NewRouter()
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/info", info).Methods("GET")

	// addr := determineListenAddress()
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	welcomeString := "Welcome to my Personal Portfolio. I'm Reljod T. Oreta, and I am a Fullstack Developer."

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(welcomeString))
}

func info(w http.ResponseWriter, r *http.Request) {

	var personalInfo PersonalInfo
	personalInfo.Name = "Reljod T. Oreta"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(personalInfo)
}

// func determineListenAddress() string {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		return ":" + "8080"
// 	}
// 	return ":" + port
// }

type PersonalInfo struct {
	Name string
}
