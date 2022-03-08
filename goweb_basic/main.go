package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//barHandler
func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "hello %s  , %d", name, id)
}
func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}
