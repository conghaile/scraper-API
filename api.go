package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) warosuHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		fmt.Println("Ayyy")
	}

	if r.Method == "POST" {
		newPost := new(bizPost)
		if err := json.NewDecoder(r.Body).Decode(newPost); err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, newPost)
	}

	return fmt.Errorf("Method not allowed: %s", r.Method)
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/warosu", makeHTTPHandlerFunc(s.warosuHandler))

	log.Println("Server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
