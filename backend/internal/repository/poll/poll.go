package poll

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type pollRepository struct {
	db *sql.DB
}

var _ repository.PollRepository = (*pollRepository)(nil)

func NewPollRepository(db *sql.DB) (*pollRepository, error) {
	const op = "poll.internal.repository.poll.NewPollRepository"

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS poll(
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
    		name TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &pollRepository{
		db: db,
	}, nil
}

func (r *pollRepository) Get(id int) (*entity.Poll, error) {
	const op = "poll.internal.repository.poll.Get"

	poll := &entity.Poll{}

	// check if exists poll
	err := r.db.QueryRow(`SELECT * FROM poll WHERE id=?`, id).Scan(&poll.ID, &poll.Name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return poll, nil
}

func (r *pollRepository) Create(name string) error {
	const op = "poll.internal.repository.poll.Save"

	stmt, err := r.db.Prepare(`
		INSERT INTO poll(name) VALUES (?)
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
