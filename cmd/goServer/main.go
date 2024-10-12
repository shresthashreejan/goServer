package main

import (
	"fmt"
	"net/http"
)

func main() {
	index := http.FileServer(http.Dir("./static"))
	http.Handle("/", index)
	http.HandleFunc("/hello", serveHTML)
	http.HandleFunc("/form", serveHTML)
	http.HandleFunc("/submit", submit)

	fmt.Printf("Server has started at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func submit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v\n", address)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("static/%s.html", r.URL.Path)
	http.ServeFile(w, r, fileName)
}
