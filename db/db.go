package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=posts password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) InitWarosu() error {
	return s.CreateWarosuTable()
}

func (s *PostgresStore) CreateWarosuTable() error {
	query := `CREATE TABLE IF NOT EXISTS Warosu (
			number serial primary key,
			text varchar(2000),
			time serial
		)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) InsertBizPost(posts *[]BizPost) error {
	query := "INSERT INTO Warosu VALUES ($1, $2, $3)"

	for _, post := range *posts {
		convertedPost := ConvertBizPost(&post)
		resp, err := s.db.Query(query, convertedPost.Number, convertedPost.Text, convertedPost.Time)

		if err != nil {
			return err
		}
		fmt.Printf("%v\n", resp)
	}
	return nil
}

func (s *PostgresStore) GetLatestPost() (*latestPost, error) {
	query := "SELECT MAX(Number) FROM Warosu"
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanPostNumber(rows)
	}

	return nil, fmt.Errorf("No posts :(")
}

func ConvertBizPost(post *BizPost) *convertedBizPost {
	return &convertedBizPost{
		Number: post.Number,
		Text:   strings.Join(post.Text, " ## "),
		Time:   post.Time,
	}
}

func scanPostNumber(rows *sql.Rows) (*latestPost, error) {
	postNumber := new(latestPost)
	err := rows.Scan(&postNumber.Number)

	return postNumber, err
}
