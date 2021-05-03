package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/api"
)

func SetupRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/heap/build", api.BuildHeap).Methods("POST")
	router.HandleFunc("/heap/sort", api.HeapSort).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s...\n", addr)
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
