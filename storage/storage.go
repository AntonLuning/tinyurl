package storage

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/AntonLuning/tiny-url/storage/sqlc"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed sqlc/schema.sql
var schema string

type Storage struct {
	db *sqlc.Queries
}

func Init(ctx context.Context, path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, err
	}

	storage := Storage{
		db: sqlc.New(db),
	}

	return &storage, nil
}

func (s *Storage) SaveURL(ctx context.Context, orignal string, shorten string) error {
	args := sqlc.CreateURLParams{
		Original: orignal,
		Shorten:  shorten,
	}

	return s.db.CreateURL(ctx, args)
}

func (s *Storage) FetchURL(ctx context.Context, shorten string) (string, error) {
	return s.db.GetOriginalURL(ctx, shorten)
}
