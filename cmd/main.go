package main

import (
	"cms"
	"net/http"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.ListenAndServe(":3000", nil)

}
