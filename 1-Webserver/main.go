package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
		
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	
	fmt.Fprintf(w, "Hello World")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	
	// if r.URL.Path != "/form" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }
	// if r.Method != "POST" {
	// 	http.Error(w, "Method is not supported.", http.StatusNotFound)
	// 	return
	// }

	fmt.Fprintf(w, "POST Successful")

	fmt.Fprintf(w, "Hello %s \n Printing Phone number %s", r.FormValue("name"), r.FormValue("phone"))
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server running on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); 
	
	err != nil {
		log.Fatal(err)
	}
}