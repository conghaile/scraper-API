package db

import "database/sql"

type BizPost struct {
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
	InsertBizPost(*[]BizPost) error
	GetLatestPost() (*latestPost, error)
}
