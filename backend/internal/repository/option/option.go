package option

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"database/sql"
	"fmt"
)

type optionRepository struct {
	db *sql.DB
}

var _ repository.OptionRepository = (*optionRepository)(nil)

func NewOptionRepository(db *sql.DB) (*optionRepository, error) {
	const op = "poll.internal.repository.option.NewOptionRepository"

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS options (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			poll_id INTEGER REFERENCES polls(id) ON DELETE CASCADE,
			name TEXT NOT NULL,
			votes INTEGER DEFAULT 0
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &optionRepository{
		db: db,
	}, nil
}

func (r *optionRepository) GetById(optionId int) (*entity.Option, error) {
	const op = "poll.internal.repository.option.GetByID"
	var option *entity.Option

	err := r.db.QueryRow(`SELECT * FROM options WHERE id = ?`, optionId).Scan(&option.ID, &option.PollID, &option.Name, &option.Votes)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return option, nil
}

func (r *optionRepository) Create(pollId int, options []entity.Option) error {
	const op = "poll.internal.repository.option.Create"

	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	// Подготавливаем SQL-запрос для вставки
	stmt, err := tx.Prepare(`
		INSERT INTO options (poll_id, name)
		VALUES (?, ?)
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	// Вставляем каждый вариант ответа
	for _, option := range options {
		_, err := stmt.Exec(pollId, option.Name)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	// Фиксируем транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *optionRepository) Vote(optionId int) error {
	const op = "poll.internal.repository.option.Vote"

	stmt, err := r.db.Prepare(`UPDATE options SET votes = votes + 1 WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(optionId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *optionRepository) GetOptionsByPollId(pollId uint) ([]*entity.Option, error) {
	const op = "poll.internal.repository.option.GetOptionsByPollID"

	stmt := `SELECT * FROM options AS o WHERE o.poll_id = ?`

	rows, err := r.db.Query(stmt, pollId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var options []*entity.Option

	for rows.Next() {
		opt := &entity.Option{}
		err := rows.Scan(&opt.ID, &opt.PollID, &opt.Name, &opt.Votes)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		options = append(options, opt)
	}

	return options, nil
}

func (r *optionRepository) GetTotalVotesByPollId(pollId uint) (uint, error) {
	const op = "poll.internal.repository.option.GetTotalVotesByPollID"

	var totalVotes uint

	err := r.db.QueryRow(`SELECT SUM(o.votes) FROM options AS o WHERE o.poll_id = ?`, pollId).Scan(&totalVotes)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return totalVotes, nil
}
