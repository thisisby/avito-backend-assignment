package postgre

import (
	"avito-backend-assignment/internal/errs"
	"avito-backend-assignment/internal/models"
	"avito-backend-assignment/internal/repositories"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type postgreLogRepository struct {
	conn *sqlx.DB
}

func NewPostgreLogRepository(conn *sqlx.DB) repositories.LogRepository {
	return &postgreLogRepository{
		conn: conn,
	}
}

func (r *postgreLogRepository) Save(log models.Log) (string, error) {
	query := `
		INSERT INTO logs(token_id, token, user_agent, url)
		VALUES (uuid_generate_v4(), $1, $2, $3)
		RETURNING token_id
	`

	var id string
	err := r.conn.QueryRow(query, log.Token, log.UserAgent, log.Url).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("postgreLogRepository.Save - r.conn.QueryRow: %w", err)
	}

	return id, nil
}

func (r *postgreLogRepository) FindById(id string) (models.Log, error) {

	tx, err := r.conn.Beginx()
	if err != nil {
		return models.Log{}, fmt.Errorf("postgreLogRepository.FindById - r.conn.Beginx: %w", err)
	}

	defer tx.Rollback()

	findByIdQuery := `
		SELECT 
			token_id,
			token,
			user_agent,
			url,
			count
		FROM logs
		WHERE token_id = $1
	`
	var log models.Log
	err = tx.Get(&log, findByIdQuery, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return models.Log{}, errs.ErrTokenNotFound
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "22P02" {
			return models.Log{}, errs.ErrInvalidUUIDFormat
		}
		return models.Log{}, fmt.Errorf("postgreLogRepository.FindById - tx.Get: %w", err)
	}

	updateCountQuery := `
		UPDATE logs
		SET count = count + 1
		WHERE token_id=$1
	`
	_, err = tx.Exec(updateCountQuery, id)
	if err != nil {
		return models.Log{}, fmt.Errorf("postgreLogRepository.FindById - tx.Exec: %w", err)
	}

	tx.Commit()

	return log, nil
}

func (r *postgreLogRepository) FindByRequestId(request string) (models.Log, error) {
	//TODO implement me
	panic("implement me")
}
