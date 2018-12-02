package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/v1/ProcessImage", ProcessImage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
	fmt.Println("Listening:8081")
}
