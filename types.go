package main

import "net/http"

type APIServer struct {
	listenAddr string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

type bizPost struct {
	Number int      `json:"number"`
	Text   []string `json:"text"`
	Time   int      `json:"time"`
}
