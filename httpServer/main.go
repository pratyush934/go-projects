package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":8080"

func greetingHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Error(w, "Sorry dear you are not on the right path", http.StatusBadRequest)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Sorry dear you are not on the right method", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Hey I am Pratyush just passing the time in the Lab by making this side project")
	fmt.Fprintf(w, "Have a good day")

}

func handleForm(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Hey you might not be making the correct request", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "You have just entered the name as %s\n", name)
	fmt.Fprintf(w, "You have just entered the address as %s\n", address)

}

func main() {

	fileSever := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileSever)
	http.HandleFunc("/home", greetingHome)
	http.HandleFunc("/form", handleForm)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
