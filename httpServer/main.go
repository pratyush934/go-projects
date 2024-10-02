package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":8080"

func greetingYou(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/home" {
		http.Error(w, "It is not the expected path that we want and you need to improve ", http.StatusBadRequest)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "It is not the expected method that we expected and you need to improve it ", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Welcome back , I am pratyush and learning to code in golang")
	fmt.Fprintln(w, "Please help me in moving my journey and want to be great as an engineer")

}

func fillTheForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		fmt.Fprintln(w, "It is not the correct path that you have chosen and we have better expectation than you")
		return
	}

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Hey %s , How are You/n", name)
	fmt.Fprintf(w, "You address is %s as you mentioned", address)
}

func main() {
	log.Println("I am trying to build a new project on golang")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/home", greetingYou)
	http.HandleFunc("/form", fillTheForm)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
