package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Good Going")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greeting", greeting).Methods("GET")
	fmt.Println("Server started at port number 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
