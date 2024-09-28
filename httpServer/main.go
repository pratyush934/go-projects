package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":8080"

func greetingMethod(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/home" {
		http.Error(w, "Sorry You might land up in wrong page", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "You are hitting wrong request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Welcome I am Pratyush")
	fmt.Println(w, "I am Pratyush and Welcome")
}

func postValue(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {

		log.Printf("Error exist and kuch nahi kiya jaa sakt hai %v", err)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Your name gentleman is %s", name)
	fmt.Fprintf(w, "Your address gentleman is %s", address)
}

func main() {

	log.Println("Http Server is running peacefully and working fine")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/home", greetingMethod)
	http.HandleFunc("/form", postValue)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
