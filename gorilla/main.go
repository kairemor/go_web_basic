package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// User type
type User struct {
	ID    string `json:"id"`
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
func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "The receive id is %s", id)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

// http.Handler
type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := User{ID: "1", Name: "Mor", Email: "Kairemor"}

	json.NewEncoder(w).Encode(u)
}

func main() {
	// http.HandleFunc("/", homeHandler)

	// http.Handle("/api", apiHandler{})

	mux := mux.NewRouter()

	mux.Handle("/home", withLogger(http.HandlerFunc(homeHandler)))
	mux.Handle("/api", withLogger(apiHandler{}))
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", userHandler).Methods("GET")

	http.ListenAndServe(":4000", mux)
}
