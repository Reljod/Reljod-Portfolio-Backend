package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/Reljod/Reljod-Portfolio-Backend/app/api"
	"github.com/Reljod/Reljod-Portfolio-Backend/config"

	// swagger embed files
	httpSwagger "github.com/swaggo/http-swagger" // gin-swagger middleware
)

const BASEPATH = config.BasePath

func SetupRouter() {
	router := mux.NewRouter()

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger", http.StatusFound)
	})
	router.HandleFunc(BASEPATH+"/heap/build", api.BuildHeap).Methods("GET", "POST")
	router.HandleFunc(BASEPATH+"/heap/sort", api.HeapSort).Methods("GET", "POST")

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
