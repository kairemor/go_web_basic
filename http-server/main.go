package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello from go server [%s]", r.URL)

	// json.NewEncoder(w).Encode("ok")
}

// http.Handler
type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from API [%s %s]", r.URL, r.Method)

}

func main() {
	http.HandleFunc("/", homeHandler)

	http.Handle("/api", apiHandler{})

	http.ListenAndServe(":4000", nil)
}
