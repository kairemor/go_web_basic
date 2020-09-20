package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello from go server [%s]", r.URL)
}

func withLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(start)
		fmt.Printf("%s %s process in %s \n", r.Method, r.URL, end)
	})
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

	mux.Handle("/home", withLogger(http.HandlerFunc(homeHandler)))
	mux.Handle("/api", withLogger(apiHandler{}))
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from User [%s %s]", r.URL, r.Method)
	})

	http.ListenAndServe(":4000", mux)
}
