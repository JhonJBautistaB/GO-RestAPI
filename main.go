package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	dbConnection()
	route := mux.NewRouter().StrictSlash(true)
	/**
	* Endpoint Route API
	**/
	route.HandleFunc("/", indexRestAPI)
	route.HandleFunc("/api/v1/matriz", getOrderMatrizView).Methods("GET")
	route.HandleFunc("/api/v1/users/", getUserView).Methods("GET")
	route.HandleFunc("/api/v1/users/", postUserView).Methods("POST")
	/**
	** Instancia del servidor e Inicialización
	**/
	server := &http.Server{
		Addr:           ":5001",
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listeng http://localhost:5001")
	log.Println(server.ListenAndServe())
}
