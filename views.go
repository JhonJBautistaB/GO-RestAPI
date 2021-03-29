package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func indexRestAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rest-API Prueba Técnica")
}

func getOrderMatrizView(w http.ResponseWriter, r *http.Request) {
	ld := []int{3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2}
	sort.Ints(ld)
	fmt.Println(ld)

}

func getUserView(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Visualizar Datos usuario desde BD SQLite3")
	n := new(users)
	response, err := n.getUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	j, err := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}

func postUserView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insertar Datos usuario a BD SQLite3")
}
