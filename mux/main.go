package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request:", r)
	return
}

func main() {

	// fmt.Println("Response Info:")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
