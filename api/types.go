package api

import (
	"net/http"

	"github.com/conghaile/simple-API/db"
)

type APIServer struct {
	listenAddr string
	store      db.Storage
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}
