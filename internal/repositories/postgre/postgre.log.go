package postgre

import (
	"avito-backend-assignment/internal/models"
	"avito-backend-assignment/internal/repositories"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type postgreLogRepository struct {
	conn *sqlx.DB
}

func NewPostgreLogRepository(conn *sqlx.DB) repositories.LogRepository {
	return &postgreLogRepository{
		conn: conn,
	}
}

func (r *postgreLogRepository) Save(log models.Log) (int, error) {
	query := `
		INSERT INTO log(id, user_agent, request_id, random_value, url, count)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := r.conn.QueryRow(query, log.ID, log.UserAgent, log.RequestId, log.RandomValue, log.Url, log.Count).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgreLogRepository.Save - r.conn.QueryRow: %w", err)
	}

	return id, nil
}

func (r *postgreLogRepository) FindById(id string) (models.Log, error) {
	query := `
		SELECT 
			id,
			user_agent,
			request_id,
			random_value,
			url,
			count,
		FROM log
		WHERE id = $1
	`
	var log models.Log
	err := r.conn.Get(&log, query, id)
	if err != nil {
		return models.Log{}, fmt.Errorf("postgreLogRepository.FindById - r.conn.Get: %w", err)
	}

	return log, nil
}

func (r *postgreLogRepository) FindByRequestId(request string) (models.Log, error) {
	//TODO implement me
	panic("implement me")
}
