package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/warosu", makeHTTPHandlerFunc(s.warosuHandler))

	log.Println("Server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) warosuHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		s.handleGetLatestPost(w, r)
	}
	if r.Method == "POST" {
		s.handleCreateWarosuPost(w, r)
	}
	return fmt.Errorf("Method not allowed: %s", r.Method)
}

func (s *APIServer) handleCreateWarosuPost(w http.ResponseWriter, r *http.Request) error {
	newPost := new(bizPost)
	if err := json.NewDecoder(r.Body).Decode(newPost); err != nil {
		return err
	}

	if err := s.store.InsertBizPost(newPost); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newPost)
}

func (s *APIServer) handleGetLatestPost(w http.ResponseWriter, r *http.Request) error {
	postnum, err := s.store.GetLatestPost()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, postnum)
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
