package main

import (
	"database/sql"
	"net/http"
)

type APIServer struct {
	listenAddr string
	store      Storage
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

type convertedBizPost struct {
	Number int
	Text   string
	Time   int
}

type PostgresStore struct {
	db *sql.DB
}

type latestPost struct {
	Number int `json:"number"`
}

type Storage interface {
	InsertBizPost(*bizPost) error
	GetLatestPost() (*latestPost, error)
}
