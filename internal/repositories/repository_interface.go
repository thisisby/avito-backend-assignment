package repositories

import "avito-backend-assignment/internal/models"

type LogRepository interface {
	Save(log models.Log) (string, error)
	FindById(id string) (models.Log, error)
	FindByRequestId(request string) (models.Log, error)
}
