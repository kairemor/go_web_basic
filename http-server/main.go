package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello from go server [%s]", r.URL)
}

// http.Handler
type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from API [%s %s]", r.URL, r.Method)
	// toujours set le content type
	w.Header().Set("Content-Type", "application/json")
	u := user{Name: "Mor", Email: "Kairemor"}

	json.NewEncoder(w).Encode(u)
}

func main() {
	// http.HandleFunc("/", homeHandler)

	// http.Handle("/api", apiHandler{})

	mux := http.NewServeMux()

	mux.Handle("/api", apiHandler{})
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from User [%s %s]", r.URL, r.Method)
	})

	http.ListenAndServe(":4000", mux)
}
