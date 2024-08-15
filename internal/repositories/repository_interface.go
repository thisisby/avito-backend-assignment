package repositories

import "avito-backend-assignment/internal/models"

type LogRepository interface {
	Save(log models.Log) (int, error)
	FindById(id string) (models.Log, error)
	FindByRequestId(request string) (models.Log, error)
}
