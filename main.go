package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main (){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}